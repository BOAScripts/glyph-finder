import os
import re
import json

base_dir = "/hme/master/git/nerd-fonts/bin/scripts/lib"
output_path = "glyphs.json"

# Files to exclude
exclude = {"i_all.sh", "i_material.sh"}

# Map filenames to human-readable group names
group_map = {
    "i_cod.sh": "Codicons",
    "i_dev.sh": "Devicons",
    "i_extra.sh": "Progress indicators",
    "i_fa.sh": "Font Awesome",
    "i_fae.sh": "Font Awesome Extension",
    "i_iec.sh": "IEC Power Symbols",
    "i_logos.sh": "Font Logos",
    "i_md.sh": "Material",
    "i_oct.sh": "Octicons",
    "i_ple.sh": "Powerline Extra",
    "i_pom.sh": "Pomicons",
    "i_seti.sh": "Seti-UI",
    "i_weather.sh": "Weather",
}

# Regex pattern: captures lines like i='' i_seti_folder=$i
pattern = re.compile(r"i='(?P<glyph>[^']+)'\s+i_[^_]+_(?P<name>[a-zA-Z0-9_]+)=\$i")

glyphs = []

for fname in os.listdir(base_dir):
    if not fname.endswith(".sh") or fname in exclude:
        continue

    group_name = group_map.get(fname, None)
    if not group_name:
        continue

    fpath = os.path.join(base_dir, fname)
    with open(fpath, "r", encoding="utf-8") as f:
        for line in f:
            match = pattern.search(line)
            if match:
                glyph = match.group("glyph")
                name = match.group("name")
                glyphs.append({"value": glyph, "name": name, "group": group_name})

# Save results to JSON
with open(output_path, "w", encoding="utf-8") as out:
    json.dump(glyphs, out, ensure_ascii=False, indent=2)

print(f"Extracted {len(glyphs)} glyphs into {output_path}")
