from aws_json_parse_withkey import get_policies_with_params
from aws_json_parse import get_policies
import pyperclip
from pathlib import Path
from subprocess import getoutput


def save_policy(file: str, policyName: str, data: str):
    path = Path.cwd() / "aws" / "policy" / f"{file}.go"
    out = f"""package policy

import "github.com/securisec/cliam/shared"

// {policyName} policy
var {policyName} = map[string]Service{{
    {data}
}}"""
    with path.open("w") as f:
        f.write(out)
    getoutput(f"go fmt {path}")


filename = "../../temp/aws-sdk-js/apis/appflow-2020-08-23.normal.json"

policies = "\n".join(get_policies(filename, True))
with_extras = "\n".join(get_policies_with_params(filename))

data = policies + "\n\n" + with_extras
# pyperclip.copy(policies + '\n\n' + with_extras)

save_policy("appflow", "AppFlowPolicies", data)
