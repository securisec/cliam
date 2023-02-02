from aws_json_parse_withkey import get_policies_with_params
from aws_json_parse import get_policies
from pathlib import Path
from subprocess import getoutput
from pyperclip import copy


def save_policy(file: str, policyName: str, data: str):
    path = Path.cwd() / "aws" / "policy" / f"{file}.go"
    out = f"""package policy

import "github.com/securisec/cliam/shared"

// {policyName}Policies policy
var {policyName}Policies = map[string]Service{{
    {data}
}}"""
    with path.open("w") as f:
        f.write(out)
    getoutput(f"go fmt {path}")


# securitylake, scheudler
filename = "../../temp/aws-sdk-js/apis/xray-2016-04-12.normal.json"
name = "xray"
policyName = "Xray"

policies = "\n".join(get_policies(filename, True))
with_extras = "\n".join(get_policies_with_params(filename))

data = policies + "\n\n" + with_extras

# copy(data)
save_policy(name, policyName, data)
