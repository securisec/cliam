from pathlib import Path
import re

extract_var = re.compile('var (\w+?) =')

paths = (Path.cwd() / "azure" / "policy").glob('*')

bad = []
for p in paths:
    try:
        d = p.read_text()
        v = extract_var.search(d).group(1)
        k = v.replace('_', '.')
        print(f'"{k}": policy.{v},')
    except:
        bad.append(p)
        continue

print(bad)