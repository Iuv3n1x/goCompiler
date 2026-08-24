# Sprachspezifikation

## 1. Variablen & Datentypen

### Deklaration & Initialisierung
* **Deklaration mit Initialwert:** `let x = 8` oder `let x = "Hallo Welt"`
* **Dynmaisch / Tipisiert:** `=` sorgt für eine Dynamische zuweisung. `:=` legt eine Typisierte Zuweisung fest. `:=` wird lediglich bei der Initalisierung eines Wertes festgelegt.
* **Deklaration ohne Initialwert:** `let x` (entspricht `let x = null`)
* **Neuzuweisung / Ändern:** Erfolgt ohne `let` (z. B. `x = 10`)

### Arrays & Typprüfung
* **Dynamisches Array:** `let x = [1, "Hallo", null, true]`
* **Typensicheres Array:** Wird ein Array mit `:=` deklariert, müssen alle Elemente denselben Datentyp besitzen:
  `let x := [1, 2, 3, 4]`

### Array-Operationen
* **Zugriff:** `x[0]` (greift auf das erste Element zu)
* **Entfernen:** `x[^0]` (entfernt das Element an Index 0)
  * `x[^0] ~>` gibt das neue Array weiter. `x[&^0]` modifiziert das Array, gibt jedoch den entfernten Wert weiter.
* **Überschreiben:** `x[0="Hallo"]` (überschreibt den Wert an Index 0 mit `"Hallo"`)
  * `x[0="Hallo"] ~>` gibt das neue Array weiter. `x[&0="Hallo"]` modifiziert das Array, gibt jedoch den neuen Wert weiter.
* **Hinzufügen:** `x[+="Hallo"]` (fügt `"Hallo"` am Ende des Arrays an)
  * `x[+="Hallo"] ~>` gibt das neue Array weiter. `x[&+="Hallo"]` modifiziert das Array, gibt jedoch den neuen Wert weiter.
* **Pipeline-Verarbeitung:** Transformationen lassen sich direkt ketten (z. B. `myArray ~> [+="Wert"] ~> [^0] ~> print`).
* **Wertvergabe mit `~|>`:** Vor dem `~|>` oder `||>` muss ein Array Index für die Wertablagerung oder `[+]` stehen.
* **Rückgabewert:** Wenn in einem `match` ein Array mit einer Aktion als Bedingung steht, dann wird geschaut ob die Aktion erfolgreich ausgeführt wurde. `array[^0]` ist solange `true`, bis kein Element zum entfernen gefunden werden kann.

---

## 2. Operatoren

### Mathematische Operatoren
* **Addition / Konkatenation:** `+` (dient auch zum Zusammensetzen von Strings)
* **Subtraktion:** `-`
* **Multiplikation:** `*`
* **Division:** `/`
* **Modulo:** `%%`

### Vergleichsoperatoren
* `<` (kleiner als)
* `>` (größer als)
* `<=` (kleiner oder gleich)
* `>=` (größer oder gleich)
* `==` (gleich)
* `!=` (ungleich)

---

## 3. Pipeline- & Flow-Operatoren (`~>` und `|>`)

### Weitergabe von Daten
* **`~>` (Asynchron / Parallel):** Reicht Daten an das nächste Glied weiter.
  * Werden mehrere Funktionen in eckigen Klammern übergeben, z. B. `[function1(), function2()] ~> exampleVariable`, werden diese **zeitgleich via Go-Routinen** ausgeführt. Ein Array aus den return Werten beider Funktionen wird weitergegeben.
  * *Best Practice:* Standardmäßig für Weitergaben nutzen.
* **`|>` (Synchron / Sequenziell):** Führt Funktionen nacheinander aus.

### Rückgabe & Steuerung
* **In-Place Modifikatoren (`~|>` und `||>`):**
  * Modifiziert die vorherige Variable direkt anort, anstatt den Wert nur weiterzureichen.
  * Beispiel: `varName ~|> len` wendet `len` direkt auf `varName` an.
  * Bei synchroner Verarbeitung entsprechend `||>`.
  * In-Place Modifikatoren sind Dynamische Operatoren. Sie können nicht mit Typisierten Elementen kombiniert werden.
* **Rückgabe via `return`:** `return` kann Ziel einer Pipeline sein (`~> return`), reicht Daten danach jedoch **nicht** weiter (`~> return ~>` ist ungültig).

---

## 4. Funktionen (`flow`)

### Syntax & Parameter
* Funktionen werden mit dem Schlüsselwort `flow` definiert:
  ```
  flow flowName(param paramType) {
      // Body
      return result
  }
  ```
* **Typisierung:** Parametertypen müssen bei der Funktionsdefinition explizit angegeben werden.

### Pipeline-Einspeisung (`param{}`)
* Soll ein Parameter direkt über eine Pipeline (`~>` / `|>`) befüllt werden, wird der Parameter mit `{}` markiert: `param{}`.
* **Überschuss-Regel:** Werden 3 Werte über eine Pipeline übergeben, aber nur 2 Parameter besitzen die `{}`-Markierung, wird der 3. Wert verworfen.

### Rückgabewerte
* Jeder `flow` **muss** ein `return` enthalten.
* Gültige Formen: `return` (leer / null), `return Wert` oder `return Variable`.

---

## 5. Eingebaute Funktionen & Keywords

### `print`
* Wird **ohne Klammern** aufgerufen.
* Reicht den empfangenen Wert transparent weiter (Pass-through).
* Beispiel: `status ~> match { ... } ~> response ~> print ~> nextStep`

### `len`
* Ermittelt die Länge eines übergebenen Werts.
* Wird **ohne Klammern** aufgerufen.
* Beispiel: `varName ~> len` oder `varName ~|> len` (speichert Ergebnis direkt in `varName`).
* Wird ein **zweiter Wert** genannt, welcher noch nicht deklariert wurde, wird dieser automatisch zum Typ **Number** und hat den Wert der Länge. Das erste Element behält seinen ursprungs Wert: `varName, varLen ~> len`

---

## 6. Control Flow (`match`)

### Pattern Matching
Bedingte Abfragen werden über den `match`-Block gesteuert:
```
let status = 200

status ~> match {
    [== 200] ~> processSuccess(),
    [default] ~> logError()
} ~> response ~> print
```
* **Default**: Die Defaultfunktion muss in jedem match Befehl existieren.

### Anwendungsbeispiele von `match`
Eine Schleife zum verarbeiten jedes Elements eines Arrays sieht wie folgt aus:
```
let items = [1, 2, 3, 4]

flow processItems(list{}) {
    list ~> match {
        [list&^0] ~> processSingleItem(param{}) ~> processItems(list),
        [default] ~> return "Alle Elemente verarbeitet"
    } ~> return
}
```

Eine Bedinungsschleife sieht wie folgt aus:
```
flow keepRunning(counter{}) {
    counter ~> match {
        [< 10] ~> (counter + 1) ~> keepRunning(param{}),
        [default] ~> counter
    } ~> return
}
```

Eine Schleife die in einem loop läuft sieht wie folgt aus:
```
flow countLoop(current{}, max{}) {
    current ~> match {
        [<= max] ~> print ~> (current + 1) ~> countLoop(param{}, max),
        [default] ~> return "Fertig"
    } ~> return
}
```

Um die Inhate eines Arrays parallel zu verarbeiten wird folgender Code verwendet:
```
myArray ~> match {
    [myArray&^0] ~> [doTaskA(param{}), doTaskB(param{})] ~> processItems(myArray),
    [default] ~> return
}
```

* **Array-Schleifen (!Kritisch!):** Arrays sollten lediglich in den geannten Schleifen fungieren, wenn sie eine bekannte Länge haben, die nicht zu groß ist. Im zweifel sollten die Funktionen `each` und `map` verwendet werden.

### `each` & `map`
`items ~> each(flow) ~>` gibt jedes Element von `items` und setzt es einzelnd in die Funktion, gibt am Ende jedoch das gesamte Array aus.

`map` setzt genau wie `each` jedes Element in ein `flow`, erstellt daraus jedoch ein neues Array, aus den Ergebnissen:
```
flow double(n{ Number }) {
    return n * 2
}

let items = [1, 2, 3, 4]

items ~> map(double) ~> print
```
