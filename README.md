# Trainingstagebuch
Funktionen:
## Training protokollieren

### Training beginnen
- Trainingseintrag wird angelegt mit Titel, Ort und Zeit 
### Übung hinzufügen
- Dem aktuellen Training wird eine Übung hinzugefügt
### Training Beenden
- Das aktuelle Training wird beendet. 

## Alle Protokollierten Trainings 
- Sortierte Übersicht über alle Trainings 
## Informationen zu  einer bestimmten Übungen abfragen
- z.B. alle Einträge für Liegestützt 
- Wann das letzte mal <Übung>
- Wieviel Wiederholungen beim letzten mal <Übung>
## Übungen eines Trainings abfragen
- z.B. alle Übungen eines bestimmten Trainings abfragen

# SQL-Sdchema?
|SessionId|Title|Location|UserId|Begin|End|


|PracticeId|SessionId|NameKey|Reps|Duration|
