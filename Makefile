.PHONY: all build generate clean test

generate:
	go generate ./...

tag:
	@echo "Creating next tag..."
	@bash -c ' \
	PREV_TAG=$$(git tag --sort=-version:refname | head -1); \
	if [ -z "$$PREV_TAG" ]; then \
		NEXT_TAG="v0.0.1"; \
	else \
		VERSION=$$(echo $$PREV_TAG | sed "s/v//"); \
		MAJOR=$$(echo $$VERSION | cut -d. -f1); \
		MINOR=$$(echo $$VERSION | cut -d. -f2); \
		PATCH=$$(echo $$VERSION | cut -d. -f3); \
		NEXT_MINOR=$$((MINOR + 1)); \
		NEXT_TAG="v$$MAJOR.$$NEXT_MINOR.0"; \
	fi; \
	echo "Previous tag: $$PREV_TAG"; \
	echo "Next tag: $$NEXT_TAG"; \
	echo ""; \
	echo "Changelog:"; \
	git changelog $$PREV_TAG..; \
	echo ""; \
	read -p "Create tag $$NEXT_TAG? [y/N] " -n 1 -r; \
	echo ""; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		CHANGELOG=$$(git changelog $$PREV_TAG..); \
		git tag -a $$NEXT_TAG -m "Version $$NEXT_TAG" -m "" -m "$$CHANGELOG"; \
		echo "Tag $$NEXT_TAG created successfully"; \
	else \
		echo "Tag creation cancelled"; \
	fi'
