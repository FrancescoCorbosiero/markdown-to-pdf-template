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

```bash
pandoc template.md -o [pdf-name].pdf --css=style.css --pdf-engine=wkhtmltopdf --resource-path=.
```

# Variabili per il Template di Documentazione WordPress

## Informazioni Generali
- `client_name`: Nome del cliente
- `document_date`: Data di creazione del documento (formato: GG/MM/AAAA)
- `doc_version`: Versione del documento (default: 1.0)

## Sezione 1: Introduzione
- `site_objectives`: Obiettivi principali del sito web
- `target_audience`: Descrizione del pubblico target
- `site_structure`: Elenco delle principali sezioni/pagine del sito

## Sezione 2: Credenziali di Accesso
### Provider del Dominio
- `domain_provider`: Nome del provider del dominio
- `domain_panel_url`: URL del pannello di controllo
- `domain_username`: Nome utente
- `domain_password`: Password (o "Fornita separatamente per motivi di sicurezza")
- `domain_notes`: Note aggiuntive sul dominio

### Hosting
- `hosting_provider`: Nome del provider di hosting
- `hosting_panel_url`: URL del pannello di controllo
- `hosting_username`: Nome utente
- `hosting_password`: Password (o "Fornita separatamente per motivi di sicurezza")
- `hosting_type`: Tipo di hosting (es. condiviso, VPS, dedicato)
- `hosting_specs`: Specifiche tecniche dell'hosting

### Webmail
- `webmail_url`: URL di accesso alla webmail
- `email_accounts`: Elenco degli indirizzi email configurati con relative funzioni
- `email_config_notes`: Note sulla configurazione email

### WordPress
- `wp_admin_url`: URL di amministrazione WordPress
- `wp_admin_username`: Nome utente amministratore
- `wp_admin_password`: Password (o "Fornita separatamente per motivi di sicurezza")
- `wp_other_users`: Elenco degli altri utenti configurati con relativi ruoli

## Sezione 3: Guide Operative
- `menu_positions`: Posizioni menu disponibili nel tema
- `backup_solution_details`: Dettagli sulla soluzione di backup implementata
- `backup_frequency`: Frequenza dei backup automatici
- `backup_storage`: Posizione di archiviazione dei backup
- `backup_content`: Dettagli sui contenuti inclusi nei backup
- `manual_backup_instructions`: Istruzioni per l'esecuzione di backup manuali

## Sezione 4: Documentazione Tecnica
### Tema
- `theme_name`: Nome del tema
- `theme_version`: Versione del tema
- `theme_type`: Tipo di tema (es. premium, personalizzato, child theme)
- `theme_developer`: Sviluppatore/agenzia
- `theme_license`: Tipo di licenza e scadenza se applicabile
- `theme_customizations`: Dettagli sulle personalizzazioni effettuate
- `custom_css_details`: Dettagli sulle modifiche CSS personalizzate
- `modified_files`: Elenco dei file modificati
- `custom_functionality`: Dettagli sulle funzionalità personalizzate
- `theme_options_path`: Percorso alle opzioni del tema
- `theme_main_options`: Elenco delle principali opzioni configurabili

### Plugin
- `essential_plugins`: Elenco dei plugin essenziali con nome, versione e scopo
- `additional_plugins`: Elenco di altri plugin con nome, versione e scopo
- `plugin_dependencies`: Eventuali dipendenze tra plugin o conflitti noti

### SEO
- `seo_plugin`: Nome del plugin SEO utilizzato
- `seo_configuration`: Dettagli sulla configurazione SEO
- `gsc_account`: Account Google Search Console
- `gsc_verification_method`: Metodo di verifica utilizzato
- `gsc_verification_date`: Data di verifica
- `ga_account`: Account Google Analytics
- `ga_tracking_id`: ID di monitoraggio
- `ga_implementation`: Metodo di implementazione
- `sitemap_url`: URL della sitemap
- `sitemap_generator`: Plugin o metodo utilizzato per generare la sitemap
- `sitemap_content`: Tipi di contenuto inclusi nella sitemap
- `sitemap_frequency`: Frequenza di aggiornamento della sitemap

### Cookie e Privacy
- `cookie_solution`: Plugin/Sistema utilizzato per la gestione dei cookie
- `cookie_configuration`: Dettagli sulla configurazione dei cookie
- `necessary_cookies`: Elenco dei cookie necessari
- `analytics_cookies`: Elenco dei cookie analitici
- `marketing_cookies`: Elenco dei cookie marketing/terze parti
- `privacy_policy_url`: URL della pagina Privacy Policy
- `cookie_policy_url`: URL della pagina Cookie Policy
- `consent_form_details`: Dettagli sul modulo di consenso
- `policies_update_date`: Data ultimo aggiornamento policies

## Sezione 5: Note Finali
- `support_company`: Nome dell'azienda di supporto
- `support_email`: Email di supporto
- `support_phone`: Telefono di supporto
- `support_hours`: Orari di disponibilità
- `support_sla`: Tempi di risposta previsti
- `initial_doc_date`: Data di creazione iniziale del documento
- `initial_author`: Autore del documento iniziale
- `changelog_entries`: Voci aggiuntive del changelog
- `agency_name`: Nome dell'agenzia/sviluppatore
- `copyright_year`: Anno del copyright