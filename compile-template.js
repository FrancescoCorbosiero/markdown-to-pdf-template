#!/usr/bin/env node

/**
 * Script per compilare il template WordPress con i dati del cliente
 * 
 * Utilizzo:
 * node compile-template.js template.md data.json output.md
 */

const fs = require('fs');
const path = require('path');

// Verifica argomenti
if (process.argv.length < 5) {
    console.error('Utilizzo: node compile-template.js template.md data.json output.md');
    process.exit(1);
}

// Percorsi dei file
const templatePath = process.argv[2];
const dataPath = process.argv[3];
const outputPath = process.argv[4];

// Leggi i file
try {
    const templateContent = fs.readFileSync(templatePath, 'utf8');
    const dataContent = fs.readFileSync(dataPath, 'utf8');
    
    // Parsing JSON
    const data = JSON.parse(dataContent);
    
    // Compila il template
    let compiledContent = templateContent;
    
    // Sostituisci tutte le variabili
    for (const [key, value] of Object.entries(data)) {
        const regex = new RegExp(`{{\\s*${key}\\s*}}`, 'g');
        compiledContent = compiledContent.replace(regex, value);
    }
    
    // Scrivi il file output
    fs.writeFileSync(outputPath, compiledContent);
    
    console.log(`Template compilato con successo: ${outputPath}`);
    
} catch (error) {
    if (error.code === 'ENOENT') {
        console.error(`Errore: File non trovato - ${error.path}`);
    } else if (error instanceof SyntaxError) {
        console.error('Errore: Il file JSON contiene errori di sintassi');
    } else {
        console.error(`Errore: ${error.message}`);
    }
    process.exit(1);
}