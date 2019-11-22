# bq table validator 
```
bq-table-validator validate <VALIDATE_YAML_PATH>
```

## Example
```
~/go/src/github.com/rerost/bq-table-validator (master) <U>
❯ cat example/validate1.yml
- name: test
  sql: |
    SELECT 1 AS id
  expect: '[{"id": 1}]'

~/go/src/github.com/rerost/bq-table-validator (master) <U>
❯ bq-table-validator validate example/validate1.yml --projectid=$GCP_PROJECT_ID

~/go/src/github.com/rerost/bq-table-validator (master) <U>
❯ cat example/validate2.yml
- name: test
  sql: |
    SELECT 1 AS id
  expect: '[{"id": 2}]'

~/go/src/github.com/rerost/bq-table-validator (master) <U>
❯ bq-table-validator validate example/validate2.yml --projectid=$GCP_PROJECT_ID
@ [0,"id"]
- 1
+ 2
```
