# git-tag-resource

***This ressource is still under construction. There is no stable release yet. Use at your own risk.***

This resource allows to manages tags in a [git](http://git-scm.com/) repository.

## Source Configuration

| Parameter       | Description                                                       | Default Value | Required/Optional |
|-----------------|-------------------------------------------------------------------|:-------------:|:-----------------:|
| `uri`           | The URI of the git repository                                     |      N/A      |     *Required*    |
| `action`        | The action to be performed on specified tag (get, new, delete)    |      new      |     *Optional*    |                                     
| `branch`        | The branch on which the tagging operation(s) will take place      |     master    |     *Optional*    |
| `private_key`   | Private key to use when interacting with the repository           |      N/A      |     *Optional*    |
| `tag_format`    | The format of the tag to be created/read/deleted/etc.             |    {SEMVER}   |     *Optional*    |
| `tag_increment` | The type of increment (major, minor, patch, num, or date)         |     major     |     *Optional*    |
| `use_date`      | The date to use (if specified in format) in the tag               |      N/A      |     *Optional*    |

### Parameters usage

#### `tag_format` possible values
* `{dated}` Date with the format YYYY.mm.DD (dot notation)
* `{dateh}` Date with the format YYYY-mm-DD (hyphen notation)
* `{SEMVER}` Version based on [semver](https://semver.org/) format
* `{RC}` Release candidate notation based on [semver](https://semver.org/spec/v2.0.0-rc.1.html) format
* `#` A positive number that starts at 1 and can then be incremented

The values are to be used to build a desired format. Here are a few examples:  
* `QA_{dateh}/v#` would be equal to `QA_2019-05-24/v1` if `QA_2019-05-24` didn't exist
* `QA_{dateh}/v#` would be equal to `QA_2019-05-24/v2` if `QA_2019-05-24` already existed with `v1`
* `{SEMVER}` would be equal to `0.0.1` if it were the first one
* `{SEMVER}-{RC}` would be equal to `0.0.1-rc.1` if it were the first one
* `{SEMVER}-{RC}` would be equal to `0.0.1-rc.2` if rc.1 already existed

When a format is given, it will be used to search all tags of the git repository matching only those with the same format.  
For example, let's say we have these tags: `0.0.1`, `0.0.2`, `0.0.3-rc.1`, `0.0.3-rc.2`  
If the format used was: `{SEMVER}-{RC}`, then the new tag would be `0.0.3-rc.3`. However, if the format used was: `{SEMVER}`, then the new tag would be `0.0.4`  
  
This allows the git-tag-resource to be used to manage QA tags, RC tags and production tags within the same pipeline, as long as they are specified separatly. 

#### `tag_increment` possible values
These only works when `tag_format` contains `{SEMVER}`.
* `major`: Updates the major version (Updates X in X.Y.Z).
* `major`: Updates the minor version (Updates Y in X.Y.Z).
* `major`: Updates the major version (Updates X in X.Y.Z).

This one works only when `tag_format` contains `#`.
* `num`: Updates the first isolated number that is not in a `{SEMVER}`, `{RC}` or date format

This one works only when `tag_format` contains either `{dated}` or `{dateh}`.
* `date`: Adds one day to the latest date found for specified pattern
    * See [use_date](#use_date-possible-values) for more details.
    
#### `use_date` possible values
***Under Construction***

### Example
Resource configuration for a repository with a private key stored in a vault and a parameter `{{RELEASE_DATE}}` stored in a yaml parameters file.
``` yaml
resources:
- name: qa-source
  type: git
  source:
    uri: ssh://git@git.abc.com:1234/test/dev/repo.git
    branch: develop
    private_key: ((git-key))
    tag_format: QA_{dateh}/v#
    tag_increment: num
    use_date: today
    
- name: prod-source
  type: git
  source:
    uri:
    branch: master
    private_keyL ((git-key))
    tag_format: PROD_{dateh}/v#
    tag_increment: num
    use_date: {{RELEASE_DATE}}
```