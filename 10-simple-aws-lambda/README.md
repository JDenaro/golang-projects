# AWS Lambda Go Function

A simple AWS Lambda function written in Go that processes user information and returns a formatted message.

## 🚀 Features

- **Go-based Lambda function** using AWS Lambda Go runtime
- **JSON input/output** handling with structured data
- **Docker support** for containerized development
- **AWS IAM integration** with proper role-based permissions

## 📋 Prerequisites

- Go 1.23.4 or higher
- AWS CLI configured with appropriate credentials
- Docker (optional, for containerized development)

## 🏗️ Project Structure

```
10-aws-lambda-go/
├── main.go              # Main Lambda function code
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
└── trust-policy.json    # IAM trust policy for Lambda role
```

## 🛠️ Installation & Setup

### 1. Install Dependencies

```bash
go mod tidy
```

### 2. Build the Lambda Binary

```bash
# Build for Linux (required for AWS Lambda)
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
```

### 3. Create Deployment Package

```bash
zip function.zip bootstrap
```

## 🚀 Deployment

### 1. Create IAM Role

```bash
# Create the trust policy
aws iam create-role \
  --role-name test-lambda \
  --assume-role-policy-document file://trust-policy.json

# Attach basic execution role
aws iam attach-role-policy \
  --role-name test-lambda \
  --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

### 2. Deploy Lambda Function

```bash
# Create the Lambda function
aws lambda create-function \
  --function-name test-go \
  --zip-file fileb://function.zip \
  --handler bootstrap \
  --runtime provided.al2 \
  --role arn:aws:iam::YOUR_ACCOUNT_ID:role/test-lambda

# Update function code (for subsequent deployments)
aws lambda update-function-code \
  --function-name test-go \
  --zip-file fileb://function.zip
```

### 3. Update Function Configuration

```bash
aws lambda update-function-configuration \
  --function-name test-go \
  --handler bootstrap \
  --runtime provided.al2
```



## 📝 Code Structure

### Main Function

```go
func HandleLambdaEvent(ctx context.Context, event MyEvent) (MyResponse, error) {
    return MyResponse{Message: fmt.Sprintf("%s is %d years old", event.Name, event.Age)}, nil
}
```

### Data Structures

```go
type MyEvent struct {
    Name string `json:"Name"`
    Age  int    `json:"Age"`
}

type MyResponse struct {
    Message string `json:"Answer:"`
}
```

## 🔧 Configuration

### Environment Variables

Set your AWS credentials:

```bash
export AWS_ACCESS_KEY_ID="your_access_key"
export AWS_SECRET_ACCESS_KEY="your_secret_key"
export AWS_SESSION_TOKEN="your_session_token"  # If using temporary credentials
export AWS_DEFAULT_REGION="us-east-1"
```

### IAM Trust Policy

The `trust-policy.json` file contains the necessary permissions for Lambda execution:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```

## 🐳 Docker Support (Optional)

For containerized development, you can create a Dockerfile:

```dockerfile
FROM public.ecr.aws/lambda/provided:al2

COPY bootstrap ${LAMBDA_TASK_ROOT}

CMD ["bootstrap"]
```

Then build and run:

```bash
# Build Docker image
docker build -t lambda-go .

# Run container
docker run -p 9000:8080 lambda-go
```

## 📊 Monitoring & Logs

### View CloudWatch Logs

```bash
aws logs describe-log-groups --log-group-name-prefix "/aws/lambda/test-go"
```

### Check Function Status

```bash
aws lambda get-function --function-name test-go
```

## 🔍 Troubleshooting

### Common Issues

1. **Runtime.InvalidEntrypoint Error**
   - Ensure binary is named `bootstrap`
   - Verify binary is compiled for Linux (`GOOS=linux`)
   - Check handler is set to `bootstrap`

2. **Permission Denied**
   - Verify IAM role has proper permissions
   - Check trust policy configuration

3. **Function Not Found**
   - Ensure function name is correct
   - Verify AWS region configuration

### Debug Commands

```bash
# Check function configuration
aws lambda get-function --function-name test-go

# View recent logs
aws logs tail /aws/lambda/test-go --follow
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test locally and on AWS
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For issues and questions:
- Check the troubleshooting section
- Review AWS Lambda documentation
- Open an issue in the repository

## 🔗 Useful Links

- [AWS Lambda Go Documentation](https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html)
- [AWS Lambda Go Runtime](https://github.com/aws/aws-lambda-go)
- [AWS CLI Documentation](https://docs.aws.amazon.com/cli/latest/userguide/)
- [Go Modules Documentation](https://golang.org/ref/mod) 