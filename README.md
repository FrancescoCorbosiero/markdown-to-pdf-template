# Markdown to PDF template

## Wordpress website comprehensive document

## Template structure

1. Document introduction & scope
2. Login credentials
   1. Domain provider
   2. Hosting service
   3. Webmail
   4. Wordpress
3. How to
4. Technical documentation
   1. Theme
   2. Plugin
   3. SEO & Google SC
   4. Cookies
5. Final notes

## Convert to PDF

Under root dir run:

pandoc template.md -o [pdf-name].pdf --css=style.css --pdf-engine=wkhtmltopdf --resource-path=.
