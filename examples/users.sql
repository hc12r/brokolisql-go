CREATE TABLE [geos] (
  [id] INTEGER PRIMARY KEY,
  [lat] TEXT,
  [lng] TEXT
);

CREATE TABLE [addresses] (
  [id] INTEGER PRIMARY KEY,
  [zipcode] TEXT,
  [city] TEXT,
  [geo_id] INTEGER,
  [street] TEXT,
  [suite] TEXT,
  FOREIGN KEY ([geo_id]) REFERENCES [geos] ([id]) ON DELETE CASCADE
);

CREATE TABLE [companies] (
  [id] INTEGER PRIMARY KEY,
  [name] TEXT,
  [bs] TEXT,
  [catchPhrase] TEXT
);

CREATE TABLE [users] (
  [id] INTEGER PRIMARY KEY,
  [name] TEXT,
  [username] TEXT,
  [email] TEXT,
  [address_id] INTEGER,
  [phone] TEXT,
  [website] TEXT,
  [company_id] INTEGER,
  FOREIGN KEY ([address_id]) REFERENCES [addresses] ([id]) ON DELETE CASCADE,
  FOREIGN KEY ([company_id]) REFERENCES [companies] ([id]) ON DELETE CASCADE
);

INSERT INTO [geos] ([id], [lat], [lng])
SELECT * FROM (
  SELECT 31, '-37.3159', '81.1496' UNION ALL
  SELECT 32, '-43.9509', '-34.4618' UNION ALL
  SELECT 33, '-68.6102', '-47.0653' UNION ALL
  SELECT 34, '29.4572', '-164.2990' UNION ALL
  SELECT 35, '-31.8129', '62.5342' UNION ALL
  SELECT 36, '-71.4197', '71.7478' UNION ALL
  SELECT 37, '24.8918', '21.8984' UNION ALL
  SELECT 38, '-14.3990', '-120.7677' UNION ALL
  SELECT 39, '24.6463', '-168.8889' UNION ALL
  SELECT 40, '-38.2386', '57.2232'
) AS source;


INSERT INTO [addresses] ([id], [zipcode], [city], [geo_id], [street], [suite])
SELECT * FROM (
  SELECT 11, '92998-3874', 'Gwenborough', 31, 'Kulas Light', 'Apt. 556' UNION ALL
  SELECT 12, '90566-7771', 'Wisokyburgh', 32, 'Victor Plains', 'Suite 879' UNION ALL
  SELECT 13, '59590-4157', 'McKenziehaven', 33, 'Douglas Extension', 'Suite 847' UNION ALL
  SELECT 14, '53919-4257', 'South Elvis', 34, 'Hoeger Mall', 'Apt. 692' UNION ALL
  SELECT 15, '33263', 'Roscoeview', 35, 'Skiles Walks', 'Suite 351' UNION ALL
  SELECT 16, '23505-1337', 'South Christy', 36, 'Norberto Crossing', 'Apt. 950' UNION ALL
  SELECT 17, '58804-1099', 'Howemouth', 37, 'Rex Trail', 'Suite 280' UNION ALL
  SELECT 18, '45169', 'Aliyaview', 38, 'Ellsworth Summit', 'Suite 729' UNION ALL
  SELECT 19, '76495-3109', 'Bartholomebury', 39, 'Dayna Park', 'Suite 449' UNION ALL
  SELECT 20, '31428-2261', 'Lebsackbury', 40, 'Kattie Turnpike', 'Suite 198'
) AS source;


INSERT INTO [companies] ([id], [name], [bs], [catchPhrase])
SELECT * FROM (
  SELECT 21, 'Romaguera-Crona', 'harness real-time e-markets', 'Multi-layered client-server neural-net' UNION ALL
  SELECT 22, 'Deckow-Crist', 'synergize scalable supply-chains', 'Proactive didactic contingency' UNION ALL
  SELECT 23, 'Romaguera-Jacobson', 'e-enable strategic applications', 'Face to face bifurcated interface' UNION ALL
  SELECT 24, 'Robel-Corkery', 'transition cutting-edge web services', 'Multi-tiered zero tolerance productivity' UNION ALL
  SELECT 25, 'Keebler LLC', 'revolutionize end-to-end systems', 'User-centric fault-tolerant solution' UNION ALL
  SELECT 26, 'Considine-Lockman', 'e-enable innovative applications', 'Synchronised bottom-line interface' UNION ALL
  SELECT 27, 'Johns Group', 'generate enterprise e-tailers', 'Configurable multimedia task-force' UNION ALL
  SELECT 28, 'Abernathy Group', 'e-enable extensible e-tailers', 'Implemented secondary concept' UNION ALL
  SELECT 29, 'Yost and Sons', 'aggregate real-time technologies', 'Switchable contextually-based project' UNION ALL
  SELECT 30, 'Hoeger LLC', 'target end-to-end models', 'Centralized empowering task-force'
) AS source;


INSERT INTO [users] ([id], [name], [username], [email], [address_id], [phone], [website], [company_id])
SELECT * FROM (
  SELECT 1, 'Leanne Graham', 'Bret', 'Sincere@april.biz', 11, '1-770-736-8031 x56442', 'hildegard.org', 21 UNION ALL
  SELECT 2, 'Ervin Howell', 'Antonette', 'Shanna@melissa.tv', 12, '010-692-6593 x09125', 'anastasia.net', 22 UNION ALL
  SELECT 3, 'Clementine Bauch', 'Samantha', 'Nathan@yesenia.net', 13, '1-463-123-4447', 'ramiro.info', 23 UNION ALL
  SELECT 4, 'Patricia Lebsack', 'Karianne', 'Julianne.OConner@kory.org', 14, '493-170-9623 x156', 'kale.biz', 24 UNION ALL
  SELECT 5, 'Chelsey Dietrich', 'Kamren', 'Lucio_Hettinger@annie.ca', 15, '(254)954-1289', 'demarco.info', 25 UNION ALL
  SELECT 6, 'Mrs. Dennis Schulist', 'Leopoldo_Corkery', 'Karley_Dach@jasper.info', 16, '1-477-935-8478 x6430', 'ola.org', 26 UNION ALL
  SELECT 7, 'Kurtis Weissnat', 'Elwyn.Skiles', 'Telly.Hoeger@billy.biz', 17, '210.067.6132', 'elvis.io', 27 UNION ALL
  SELECT 8, 'Nicholas Runolfsdottir V', 'Maxime_Nienow', 'Sherwood@rosamond.me', 18, '586.493.6943 x140', 'jacynthe.com', 28 UNION ALL
  SELECT 9, 'Glenna Reichert', 'Delphine', 'Chaim_McDermott@dana.io', 19, '(775)976-6794 x41206', 'conrad.com', 29 UNION ALL
  SELECT 10, 'Clementina DuBuque', 'Moriah.Stanton', 'Rey.Padberg@karina.biz', 20, '024-648-3804', 'ambrose.net', 30
) AS source;


