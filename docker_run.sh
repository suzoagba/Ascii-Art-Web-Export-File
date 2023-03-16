echo "Building Docker image"
docker image build -t image1 .
echo "Running Docker container"
docker run -dp 8080:8080 --name container1 image1
echo "Container running on port 8080"
echo "http://localhost:8080/"