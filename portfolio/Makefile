openapi:
	cd backend/gateway && oapi-codegen -generate types,server,spec,strict-server -package gateway -o server.gen.go ../../api/openapi/openapi.yaml

upload: openapi
	aws s3 cp ./api/openapi/openapi.yaml s3://coplay-test/portfolio.yaml --acl public-read
	portman -l ./api/openapi/openapi.yaml -b http://localhost:3000 --syncPostman --postmanFastSync