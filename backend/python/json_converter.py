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

# Write formatted data to output file
with open('proverbs_cleaned.json', 'w') as f:
    json.dump(formatted_data, f, indent=4)  # Write JSON with indentation for readability
