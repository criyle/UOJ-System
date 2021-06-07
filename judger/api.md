# UOJ judger protocol

## judge_client

### Configuration

In `.conf.json`:

- `uoj_protocol`, `uoj_host`: (http, host)
- `judger_name`:
- `judger_password`:
- `socket_port`:
- `socket_password`:

### Command line arguments

- (no parameter): start_judger_server
- start: starts as daemon, stderr -> `/log/judge.log`
- update: send cmd update with password to localhost with tcp
- stop: send cmd stop with password to localhost with tcp

### Overall

After initialization, the judger starts socket_server and judger_loop

#### API

Credential is `judger_name` and `password` (`judger_password`) json data.

Download `uri` makes `POST` request to `/judge/download/:uri` with Credential.

Submit makes `POST` request to `judge/submit` with data, credential and files, it also gets new submission if `fetch_new` is not `false`.

#### socket_server

listening tcp on `judger_port`, task are send to `taskQ` and execute main thread:

- requires json, checks `password` with `socket_password`
- cmd `stop`
  - wait on judging task finish and quit
  - writes "ok" as response
- cmd `update`
  - download `/judger` as `judger_update.zip`, `unzip` it and compile
  - if `main_judger`, post `/judge/sync-judge-client` with credential
    - expects "ok" as response
  - restarts by `execl`

#### judger_loop

Submission

```typescript
interface Submission {
    id: string;
    problem_id: string;
    problem_mtime: string;
    content: {
        file_name: string
        config: { [k:string]: string};
    };
    is_hack?: boolean;
    hack: {
        id: string;
        input_type: string; // 'USE_FORMATTER'
        input: string;
    }
    is_custom_test?: boolean;
};
```

- poll wait for new submission (2s interval)
- update problem data
  - check `problem_mtime` with `uoj_judger/data/:problem_id`
  - if `problem_mtime` > folder `mtime`
    - delete older problem data (preserving 100)
    - download `problem/:problem_id` as `:problem_id.zip` unzip and put into `data` path
- download `content.file_name` as `work/all.zip` and `unzip`
- write config `content.config` to `submission.conf`
- if `is_hack`
  - if `hack.input_type` == "USE_FORMATTER":
    - download `hack.input` to `work/hack_input_raw.txt`
    - run `run/formatter` input `work/hack_input_raw.txt` output `work/hack_input.txt`
  - else:
    - download `hack.input` to `work/hack_input_raw.txt`
  - if `is_custom_test`, write "custom_test on" to `submission.conf`
  - starts report_loop in another thread
  - run `main_judger`
  - get result from `result/result.txt`

```typescript
interface Result{
    id: string; // :hack.id or :id
    submit: boolean; // true
    is_hack?: boolean; // :is_hack
    is_custom_test?: boolean; // :is_custom_test
    result: {
        score: number;
        time: number;
        memory: number;
        details: string; // lines after details
        error: string; // content after error
        [k: string]: string; // and other key-value pairs
    }
}
```

if `is_hack`, submit file `hack_input` from `work/hack_input.txt` and `std_output` from `work/std_output.txt`

Report Loop:

- 0.2s polling rate
  - read `/result/cur_status.txt` with lock
  - submit with updated status with credential

```json
{
    "update-status": true,
    "id": ":id",
    "is_custom_test": ":is_custom_test",
    "status": "status"
}
```

## main_judge

- read conf from `work/submission.conf` and `data/:problem_id/problem.conf`.
- copy `data/:problem_id/require/*` to `work`
- if `use_builtin_judger on` then set `judger` to `builtin/judger/judger`, else `data/:problem_id/judger`
- result file set to `result/run_judger_result.txt`

## builtin/judger

- includes `uoj_judger.h`. uses `run_program` to compile and run
