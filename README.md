## JSON

Esempi di json decode/encode, azioni chiamate rispettivamente marshal e unmarshal.
Presente inoltre chiamata remota ad API Github per il recupero di dato in formato JSON 
e mapping in varie strutture. È presente funzionalità per rendering in template html.

# Argomenti trattati
- Json marshall/unmarshall ( endoce/decode )
- Chiamata a web service JSON
- Rendering JSON formattato
- Rendering in template HTML ( sintassi di Angular, anche con pipe )
- Savataggio output in locale

# Utilizzo

Compilare

```
go build
```

Eseguire e salvare output in file html
```
./json repo:golang/go is:open json decoder > test.html
```
