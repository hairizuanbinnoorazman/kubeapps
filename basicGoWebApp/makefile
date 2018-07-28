build:
	docker build -t localhost:5000/web:$(version) .
	docker push localhost:5000/web:$(version)
init:
	kubectl run web --image localhost:5000/web:$(version)
	kubectl expose --port 8080 --type NodePort deployments web
update:
	kubectl set image deployment/web web=localhost:5000/web:$(version)
cleanup:
	kubectl delete services web
	kubectl delete deployments web
