CREATE TABLE Medstores (
    ID SERIAL NOT NULL PRIMARY KEY,
    Address VARCHAR,
    Name VARCHAR
);

CREATE TABLE Drugs (
    ID SERIAL NOT NULL PRIMARY KEY,
    Name VARCHAR,
    Effects VARCHAR,
    NeedsReceipt BOOLEAN,
    Dosage VARCHAR,
    UnitWeight INT,
    MinAge INT,
    Description VARCHAR
);

CREATE TABLE DrugMap (
    ID SERIAL NOT NULL PRIMARY KEY,
    DrugId SERIAL references Drugs(ID),
    MedstoreId SERIAL references Medstores(ID),
    Price INT
);

CREATE TABLE Components (
    ID SERIAL NOT NULL PRIMARY KEY,
    Name VARCHAR
);

CREATE TABLE DrugsComponents (
    ID SERIAL NOT NULL PRIMARY KEY,
    DrugId SERIAL references Drugs(ID),
    ComponentId SERIAL references Components(ID)
);

CREATE TABLE Indications (
    ID SERIAL NOT NULL PRIMARY KEY,
    Name VARCHAR NOT NULL
);

-- Создание промежуточной таблицы для связи "Многие ко многим" между лекарствами и показаниями
CREATE TABLE DrugIndications (
    DrugId SERIAL NOT NULL references Drugs(ID),
    IndicationId SERIAL NOT NULL references Indications(ID),
    PRIMARY KEY (DrugId, IndicationId)
);

INSERT INTO Medstores (Address, Name) VALUES 
('abc', 'Лекарства 24/7'),
('bcd', 'Электронный Знахарь'),
('efg', 'Лекарь'),
('efg', 'Аптека СМ'),
('efg', 'Вита Про'),
('efg', 'Фарма Авиценна'),
('efg', 'Айболит');
INSERT INTO Drugs (Name, Effects, NeedsReceipt, Dosage, UnitWeight, MinAge, Description) VALUES 
('Парацетамол', 'Снижение боли и температуры', FALSE, '500мг', 500, 6, 'Обезболивающее и жаропонижающее средство'),
('Ибупрофен', 'Обезболивание, противовоспалительное', FALSE, '200мг', 200, 12, 'Нестероидное противовоспалительное средство'),
('Амоксициллин', 'Антибактериальная терапия', TRUE, '250мг', 250, 0, 'Антибиотик для лечения бактериальных инфекций'),
('Цитрамон', 'Снятие головной боли', FALSE, '300мг', 300, 12, 'Комбинированное средство от головной боли'),
('Аспирин', 'Снижение температуры, разжижение крови', FALSE, '100мг', 100, 16, 'Ацетилсалициловая кислота, используется для разжижения крови'),
('Кларитин', 'Лечение аллергии', FALSE, '10мг', 10, 6, 'Антигистаминный препарат'),
('Лоперамид', 'Остановка диареи', FALSE, '2мг', 2, 12, 'Средство от диареи'),
('Цефтриаксон', 'Антибактериальная терапия', TRUE, '1г', 1000, 0, 'Антибиотик для лечения инфекций'),
('Но-шпа', 'Снятие спазмов', FALSE, '40мг', 40, 6, 'Спазмолитик'),
('Омепразол', 'Снижение кислотности', FALSE, '20мг', 20, 18, 'Средство от изжоги и гастрита'),
('Нурофен', 'Обезболивание, противовоспалительное', FALSE, '400мг', 400, 12, 'Средство для снятия воспаления и боли'),
('Супрастин', 'Лечение аллергии', FALSE, '25мг', 25, 3, 'Антигистаминный препарат'),
('Мезим', 'Улучшение пищеварения', FALSE, '125мг', 125, 0, 'Ферментный препарат для улучшения пищеварения'),
('Диклофенак', 'Обезболивание, противовоспалительное', FALSE, '50мг', 50, 18, 'Средство для снятия боли и воспаления'),
('Эналаприл', 'Снижение давления', TRUE, '5мг', 5, 18, 'Препарат для лечения гипертонии'),
('Кеторол', 'Обезболивание', FALSE, '10мг', 10, 16, 'Сильное обезболивающее средство'),
('Азитромицин', 'Антибактериальная терапия', TRUE, '500мг', 500, 0, 'Антибиотик широкого спектра действия'),
('Фестал', 'Улучшение пищеварения', FALSE, '150мг', 150, 0, 'Ферментный препарат для улучшения работы ЖКТ'),
('Цитрин', 'Лечение аллергии', FALSE, '10мг', 10, 6, 'Антигистаминный препарат нового поколения'),
('Лазолван', 'Улучшение отхождения мокроты', FALSE, '30мг', 30, 0, 'Препарат для лечения кашля и мокроты');

-- Заполнение таблицы Components (Компоненты)
INSERT INTO Components (Name) VALUES 
('Парацетамол'),
('Ибупрофен'),
('Амоксициллин'),
('Кофеин'),
('Ацетилсалициловая кислота'),
('Лоратадин'),
('Лоперамид'),
('Цефтриаксон'),
('Дротаверин'),
('Омепразол'),
('Ферменты'),
('Азитромицин'),
('Диклофенак'),
('Кеторолак'),
('Сахароза'),
('Крахмал'),
('Лактоза'),
('Мукосахариды'),
('Хлоропирамин'),
('Бромгексин');

-- Заполнение таблицы DrugsComponents (Компоненты лекарств)
INSERT INTO DrugsComponents (DrugId, ComponentId) VALUES 
(1, 1), -- Парацетамол содержит компонент Парацетамол
(2, 2), -- Ибупрофен содержит компонент Ибупрофен
(3, 3), -- Амоксициллин содержит компонент Амоксициллин
(4, 1), -- Цитрамон содержит Парацетамол
(4, 4), -- Цитрамон содержит Кофеин
(4, 5), -- Цитрамон содержит Ацетилсалициловую кислоту
(5, 5), -- Аспирин содержит Ацетилсалициловую кислоту
(6, 6), -- Кларитин содержит Лоратадин
(7, 7), -- Лоперамид содержит Лоперамид
(8, 8), -- Цефтриаксон содержит Цефтриаксон
(9, 9), -- Но-шпа содержит Дротаверин
(10, 10), -- Омепразол содержит Омепразол
(11, 2), -- Нурофен содержит Ибупрофен
(12, 19), -- Супрастин содержит Хлоропирамин
(13, 11), -- Мезим содержит Ферменты
(14, 13), -- Диклофенак содержит Диклофенак
(15, 10), -- Эналаприл содержит Омепразол
(16, 14), -- Кеторол содержит Кеторолак
(17, 12), -- Азитромицин содержит Азитромицин
(18, 11), -- Фестал содержит Ферменты
(19, 6), -- Цитрин содержит Лоратадин
(20, 20); -- Лазолван содержит Бромгексин

-- Заполнение таблицы DrugMap (Доступность лекарств в аптеках)
INSERT INTO DrugMap (DrugId, MedstoreId, Price) VALUES 
(1, 1, 50), (1, 2, 55), (1, 4, 60), -- Парацетамол в 3 аптеках
(2, 2, 100), (2, 3, 105), (2, 5, 95), -- Ибупрофен в 3 аптеках
(3, 3, 200), (3, 4, 210), (3, 6, 190), -- Амоксициллин в 3 аптеках
(4, 1, 70), (4, 4, 75), (4, 7, 65), -- Цитрамон в 3 аптеках
(5, 2, 60), (5, 5, 58), (5, 6, 62), -- Аспирин в 3 аптеках
(6, 4, 150), (6, 6, 140), (6, 7, 145), -- Кларитин в 3 аптеках
(7, 5, 90), (7, 2, 95), (7, 3, 85), -- Лоперамид в 3 аптеках
(8, 3, 300), (8, 4, 320), (8, 7, 290), -- Цефтриаксон в 3 аптеках
(9, 1, 120), (9, 6, 125), (9, 7, 115), -- Но-шпа в 3 аптеках
(10, 4, 250), (10, 2, 240), (10, 5, 230), -- Омепразол в 3 аптеках
(11, 1, 140), (11, 3, 135), (11, 7, 145), -- Нурофен в 3 аптеках
(12, 6, 130), (12, 4, 125), (12, 7, 120), -- Супрастин в 3 аптеках
(13, 5, 180), (13, 1, 190), (13, 7, 170), -- Мезим в 3 аптеках
(14, 2, 200), (14, 5, 210), (14, 6, 190), -- Диклофенак в 3 аптеках
(15, 7, 250), (15, 1, 240), (15, 4, 230), -- Эналаприл в 3 аптеках
(16, 3, 300), (16, 5, 280), (16, 7, 270), -- Кеторол в 3 аптеках
(17, 1, 400), (17, 2, 410), (17, 3, 420), -- Азитромицин в 3 аптеках
(18, 4, 220), (18, 5, 210), (18, 7, 230), -- Фестал в 3 аптеках
(19, 6, 170), (19, 1, 160), (19, 5, 150), -- Цитрин в 3 аптеках
(20, 7, 190), (20, 2, 180), (20, 3, 200); -- Лазолван в 3 аптеках


INSERT INTO Indications (Name) VALUES 
('Головная боль'),
('Лихорадка'),
('Аллергия'),
('Бактериальные инфекции'),
('Диарея'),
('Спазмы'),
('Изжога'),
('Воспаление'),
('Гипертония'),
('Обезболивание'),
('Тошнота'),
('Пищеварение'),
('Разжижение крови'),
('Заболевания дыхательных путей');

-- Заполнение таблицы "DrugIndications" (Связь лекарств с показаниями)
INSERT INTO DrugIndications (DrugId, IndicationId) VALUES 
(1, 1), -- Парацетамол - Головная боль
(1, 2), -- Парацетамол - Лихорадка
(2, 1), -- Ибупрофен - Головная боль
(2, 2), -- Ибупрофен - Лихорадка
(3, 4), -- Амоксициллин - Бактериальные инфекции
(4, 1), -- Цитрамон - Головная боль
(4, 2), -- Цитрамон - Лихорадка
(5, 2), -- Аспирин - Лихорадка
(6, 3), -- Кларитин - Аллергия
(7, 5), -- Лоперамид - Диарея
(8, 4), -- Цефтриаксон - Бактериальные инфекции
(9, 6), -- Но-шпа - Спазмы
(10, 7), -- Омепразол - Изжога
(11, 1), -- Нурофен - Головная боль
(12, 3), -- Супрастин - Аллергия
(13, 12), -- Мезим - Пищеварение
(14, 8), -- Диклофенак - Воспаление
(15, 9), -- Эналаприл - Гипертония
(16, 10), -- Кеторол - Обезболивание
(17, 4), -- Азитромицин - Бактериальные инфекции
(18, 12), -- Фестал - Пищеварение
(19, 3), -- Цитрин - Аллергия
(20, 13); -- Лазолван - Заболевания дыхательных путей