import json
import os



# Read data from input file
with open('proverbs_data.json', 'r') as f:
    data = f.readlines()

# Process each line of data
formatted_data = []
for line in data:
    # Load JSON from each line
    try:
        json_data = json.loads(line)
        formatted_data.append(json_data)
    except json.JSONDecodeError as e:
        print(f"Error decoding JSON: {e}")

proverbs = json_data["Proverbs"]

# convert to list of proverbs
proverbs_list = [proverbs[key] for key in proverbs]

# create new JSON object
proverbs_json ={
    "Proverbs" : proverbs_list
}
print(proverbs_json)

# Write json data to file
with open('proverbs_only.json', 'w') as f:
    json.dump(proverbs_json, f, indent=4)

# Write formatted data to output file
with open('proverbs_cleaned.json', 'w') as f:
    json.dump(formatted_data, f, indent=4)  # Write JSON with indentation for readability