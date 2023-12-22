.PHONY: watch-css
watch-css:
	npx tailwindcss -i ./public/styles.css -o ./public/dist/css/tailwind.css --watch

