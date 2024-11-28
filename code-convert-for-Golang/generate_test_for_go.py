import os
import llmClient
import fileAccess
import time
from typing import List
from dotenv import load_dotenv


# Load environment variables from .env file
load_dotenv()

# Add a global variable to store the progress
progress = 0
file_list = []


def generate_new_test(path, prompt):
    global progress
    # Reset the progress
    progress = 0

    # Initialize the LLM client
    client = llmClient.init()

    # Get all Go files in the specified directory
    go_files = [f for f in os.listdir(path) if f.endswith('.go')]
    total_files = len(go_files)

    # Create a directory for the generated test files
    test_path_dir = os.path.join(path, "generated_tests")
    if not os.path.exists(test_path_dir):
        os.makedirs(test_path_dir)
    j = 0
    # Iterate over each Go file to generate and run tests
    for filename in go_files:
        fullfilepath = os.path.join(path, filename)
        test_filename = filename.replace(".go", "_test.go")
        test_filepath = os.path.join(test_path_dir, test_filename)
        # update the file list
        file_list.append({"name": filename, "status": "Processing"})

        # Load the content of the Go file
        filecontent = fileAccess.load_file(fullfilepath)

        # Prepare the prompt for generating test cases
        test_prompt = prompt + "\nGenerate test cases for the following Go code:\n" + filecontent

        # Execute the prompt using the LLM client
        response = llmClient.execute_prompt(
            test_prompt, "generate Go test cases", client)
        test_code = llmClient.process_llm_response(response)

        # Output the generated test code to a new file
        fileAccess.output_content(test_filepath, test_code)

        # update the file list
        for file in file_list:
            if file["name"] == filename:
                file["status"] = "Test File Generated"
                break
        time.sleep(3)
        # update the file list
        for file in file_list:
            if file["name"] == filename:
                file["status"] = "Start to run test"
                break
        # Run the generated test file using 'go test'
        os.system(f"go test {test_filepath}")

        # update the file list
        for file in file_list:
            if file["name"] == filename:
                file["status"] = "Test Done"
                break

        # update the progress
        progress = (j + 1) / total_files
        print(f"Processing file {j + 1} of {total_files}")
        j = j+1
        print(f"current progress: {progress}")
