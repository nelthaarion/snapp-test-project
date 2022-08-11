echo Building publisher app...;
cd ../Publisher && env GOOS=linux CGO_ENABLED=0 go build -o app ./cmd/api;
echo Building publisher app finished!;

echo Building consumer app...;
cd ../Consumer && env GOOS=linux CGO_ENABLED=0 go build -o app ./cmd/api;
echo Building consumer app finished!;

echo Building containers started...;
cd ../Workspace && docker-compose up -d --build 
echo Building containers done!;

