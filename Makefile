VERSION ?= $(patsubst v%,%,$(shell git describe))

bin/pulumi-resource-fivetran: provider/cmd/pulumi-resource-fivetran/schema.json
	go build -o bin/pulumi-resource-fivetran ./provider/cmd/pulumi-resource-fivetran

bin/pulumi-tfgen-fivetran: provider/cmd/pulumi-tfgen-fivetran/*.go go.sum provider/*.go
	go build -o bin/pulumi-tfgen-fivetran ./provider/cmd/pulumi-tfgen-fivetran

provider/cmd/pulumi-resource-fivetran/schema.json: bin/pulumi-tfgen-fivetran
	bin/pulumi-tfgen-fivetran $(VERSION) schema --out ./provider/cmd/pulumi-resource-fivetran

schema: provider/cmd/pulumi-resource-fivetran/schema.json

python-sdk: provider/cmd/pulumi-resource-fivetran/schema.json
	rm -rf sdk
	bin/pulumi-tfgen-fivetran $(VERSION) python
	cp README.md sdk/python/
	cd sdk/python/ && \
		sed -i.bak -e "s/0\.0\.0/$(VERSION)/g" setup.py && \
		rm setup.py.bak
