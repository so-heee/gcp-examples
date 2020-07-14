RUNTIME := go113
PROJECT_ID_DEV := recotori-dev-61c27

.PHONY: deploy-dev
deploy-dev:
	gcloud functions deploy HelloPubSub \
	--trigger-topic helloworld \
	--runtime $(RUNTIME) \
	--env-vars-file ./config/.env.develop.yaml \
	--project $(PROJECT_ID_DEV)