RUNTIME := go113
PROJECT_ID_DEV := 

.PHONY: deploy
deploy-dev:
	gcloud functions deploy HelloPubSub \
	--trigger-topic helloworld \
	--runtime $(RUNTIME) \
	--env-vars-file ./config/.env.develop.yaml \
	--project $(PROJECT_ID_DEV)