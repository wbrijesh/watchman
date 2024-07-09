CREATE TABLE Projects (
  ID TEXT PRIMARY KEY,
  Name TEXT NOT NULL
);
CREATE TABLE Logs (
  Time INTEGER NOT NULL,
  Level TEXT NOT NULL,
  Message TEXT NOT NULL,
  Subject TEXT,
  UserID TEXT,
  ProjectID TEXT NOT NULL,
  FOREIGN KEY (ProjectID) REFERENCES Project(ID)
);
CREATE TABLE ProjectKeys (
    ProjectID TEXT NOT NULL,
    AccessKey TEXT NOT NULL,
    SecretKey TEXT NOT NULL,
    Expiry INTEGER,
    FOREIGN KEY (ProjectID) REFERENCES Projects(ID)
);
