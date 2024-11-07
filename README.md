# **go-aws-lambda**
## **Project Overview**
The go-aws-lambda project provides a straightforward demonstration of creating an AWS Lambda function using the Go programming language. The function handles a custom event and performs operations based on that event in an AWS environment. The project is structured to facilitate easy deployment to AWS Lambda using the Go custom runtime.

## **Project Structure**
```bash
go-aws-lambda/
├── bootstrap               # Custom Lambda runtime script to invoke Go functions
├── function.zip            # Zipped Lambda function archive (includes Go binary)
├── go.mod                  # Go module file for dependency management
├── go.sum                  # Go checksum file for module integrity
├── main.go                 # The Go function logic executed by Lambda
├── output.txt              # Sample output log from Lambda execution
├── trust-policy.json       # IAM trust policy defining Lambda permissions
└── README.md               # Project documentation
```
### **Explanation of Files**
- **main.go**
The core Lambda function written in Go. This file contains the handler function that is executed when an event triggers the Lambda. It processes the event and returns a response.

- **go.mod & go.sum**
These files define and manage the dependencies required to build the Go Lambda function. They ensure that the necessary libraries are correctly installed and that their versions are locked for reproducibility.

- **trust-policy.json**
This file contains the IAM trust policy required to grant permissions to the Lambda function. It defines who can invoke the Lambda function. Typically, this policy would allow AWS services like API Gateway to invoke the Lambda.

- **function.zip**
This file contains the Go binary, along with any dependencies, bundled into a ZIP file. This archive is uploaded to AWS Lambda for deployment.

- **bootstrap**
The bootstrap file is the entry point for the Lambda function. When using Go with AWS Lambda, a custom runtime is required, and the bootstrap file is the standard mechanism for invoking the Go binary in this environment.

- **output.txt**
A log file capturing the output from the Lambda execution. It serves as an example of how Lambda processes an event and produces output.

## **How to Set Up and Use the Lambda Function**
### **1. Create a Lambda Function in AWS**
- Log into your AWS account and navigate to the Lambda service.
- Create a new Lambda function with a custom runtime.
- Use the bootstrap file as the entry point for the function.
### **2. Prepare the Function for Deployment**
- First, compile the Go code to generate the binary for your Lambda function:

```bash
GOOS=linux GOARCH=amd64 go build -o main main.go
```
- Once the binary is compiled, create a ZIP archive of the binary and the necessary dependencies:

```bash
zip function.zip main
```
- Upload the function.zip to AWS Lambda.

### **3. IAM Roles and Permissions**
- Configure the trust-policy.json to grant the necessary permissions for your Lambda function.
- Attach this policy to the Lambda execution role, ensuring it has permissions to access resources required for execution.
### **4. Test the Lambda Function**
- Test the function by invoking it either manually from the AWS Lambda console or via an event trigger (e.g., API Gateway or CloudWatch).
- You can monitor the Lambda execution through AWS CloudWatch logs, where you can view output similar to what's found in output.txt.
### **5. Update or Modify the Function**
- Modify main.go to change the Lambda function's behavior. After making changes, recompile and re-zip the project, then upload the new function.zip to AWS Lambda.
## **Conclusion**
This project serves as a simple example for deploying a Go-based Lambda function to AWS. It demonstrates the use of the custom runtime with Go and provides the necessary files for packaging and deploying the function. By following the steps in this documentation, developers can easily deploy Go-based Lambda functions, modify them, and integrate them with other AWS services.