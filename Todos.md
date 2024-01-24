# MNR BRIG

## Bücher
### Datenbank

#### books
| book_id | book_title | book_price | updated_at | created_at | deleted_at |
|---------|------------|------------|------------|------------|------------|
| 1 | Scoffield Bible | 45.00 | 2024-01-023 15:01:16 | 2024-01-023 15:01:16 |
| 2 | New Buch | 120.00 | 2024-01-023 15:40:14 | 2024-01-023 15:01:16 |

```sql
    SELECT * FROM books
    WHERE deleted_at IS NULL;

    -- deleted_at
    UPDATE books 
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE book_id = 1;

    UPDATE books
    SET book_title="neues Buch", book_price="20.30"
    WHERE book_id = 1;

    INSERT INTO books (book_title, book_price)
    VALUE ("Der Buchtitel", 12.35)


```
### Ablauf Books
- Erfassen von einem neuen Buch
    - ~~Speichern in der DB~~
    - Kontrolle das der Preis eine Zahl ist
    - Kontrolle das der Titel < 255 zeichen ist
    - Aktualisieren der Liste mit Bücher
- Löschen von einem Buch
    - Anklicken Löschbutton
    - UPDATE Statment absetzen und deleted_at auf aktuelle Zeit setzen
    - Tabelle aktualisiseren
- Ändern eines Bucher
    - Auswählen eines Buches
    - Ändern des DS
    - Kontrolle vor dem Speichern wie beim Erfassen
    - UDATE auf DB ausführen
    - Tabelle aktualisieren


#### sales
| sale_id | book_id | sale_date | sale-desc | updated_at | created_at |
|---------|--------|------------|------------|------------|------------|
| 1 | 1 | 23.01.24 12:00 | Verkauft Sonntag | 2024-01-023 15:01:16 | 2024-01-023 15:01:16 |
| 2 | 1 | 23.01.24 13:00 | Verkauft Sonntag | 2024-01-023 15:01:16 | 2024-01-023 15:01:16 |
```sql
    SELECT book_title, book_price, sale_date
    FROM sales AS s    
    INNER JOIN books AS b ON s.book_id = b.book_id
    ORDER BY sale_date;
```

## Spenden

#### donates
| donate_id | donate_date | donate_price | donate-desc | updated_at | created_at |
|---------|-----------|------------|------------|------------|------------|
| 1 | 23.01.24 12:00 | 250 | Verkauft Sonntag | 2024-01-023 15:01:16 | 2024-01-023 15:01:16 |
| 2 | 23.01.24 13:00 | 100 | Verkauft Sonntag | 2024-01-023 15:01:16 | 2024-01-023 15:01:16 |

```sql

SELECT * FROM sales;

```

