-- Drop existing tables (if necessary)
DROP TABLE IF EXISTS User;
DROP TABLE IF EXISTS Appointment;
DROP TABLE IF EXISTS Test;
DROP TABLE IF EXISTS TestTrack;
DROP TABLE IF EXISTS TestResults;
DROP TABLE IF EXISTS Medicine;
DROP TABLE IF EXISTS prescription;
DROP TABLE IF EXISTS vacation;
DROP TABLE IF EXISTS disease;
DROP TABLE IF EXISTS paitentDesies;



-- Create the User table
CREATE TABLE User (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    UserId INTEGER NOT NULL,
    first_name VARCHAR(16) NOT NULL,
    last_name VARCHAR(16) NOT NULL,
    Gender VARCHAR(16) NOT NULL,
    date_of_birth DATE,
    email VARCHAR(30) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    Rule CHAR(60) NOT NULL
);

-- Create the Apointments table
CREATE TABLE Appointment (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    Doctor_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    date_of_Appointment DATE,
    TheTime TIME,
    FOREIGN KEY (Doctor_id) REFERENCES User(id),
    FOREIGN KEY (user_id) REFERENCES User(id)
);

-- Create the Tets table
CREATE TABLE Test (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    Doctor_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    lab_Tech INTEGER NOT NULL,
    Test_Type VARCHAR(30) NOT NULL,
    Test_Date TIME,
    FOREIGN KEY (Doctor_id) REFERENCES User(id),
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (lab_Tech) REFERENCES User(id)

);

-- Create the TetsResults table
CREATE TABLE TestTrack (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    Test_id INTEGER NOT NULL,
    test_Statues text,
    FOREIGN KEY (Test_id) REFERENCES Tets(id)
);

-- Create the Medicine table
CREATE TABLE Medicine (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    MedicineName VARCHAR(30) NOT NULL,
    Medicine_description text
);

-- Create the prescription table
CREATE TABLE prescription (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    Doctor_id INTEGER NOT NULL,
    patient_Id INTEGER NOT NULL,
    Medicine_id INTEGER NOT NULL,
    date_of_prescription DATE,
    prescription TEXT NOT NULL,
    FOREIGN KEY (Doctor_id) REFERENCES User(id),
    FOREIGN KEY (patient_Id) REFERENCES User(id),
    FOREIGN KEY (Medicine_id) REFERENCES Medicine(id)
);


-- Create the prescription table
CREATE TABLE TestResults (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    Test_Id INTEGER NOT NULL,
    Result TEXT NOT NULL,
    FOREIGN KEY (Test_Id) REFERENCES Tets(id)
);

-- Create the prescription table
CREATE TABLE vacation (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    Doctor_id INTEGER NOT NULL,
    patient_Id INTEGER NOT NULL,
    Number_Of_Days INTEGER NOT NULL,
    Start_from DATE,
    VacationReason TEXT NOT NULL,
    FOREIGN KEY (Doctor_id) REFERENCES User(id)
    FOREIGN KEY (patient_Id) REFERENCES User(id)

);

-- Create the Medicine table
CREATE TABLE diseases (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    diseaseName VARCHAR(30) NOT NULL,
    disease_description text,
    disease_Medicine_id int,
        FOREIGN KEY (disease_Medicine_id) REFERENCES Medicine(id)

);


-- Create the Medicine table
CREATE TABLE paitentDesies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    patient_Id INTEGER NOT NULL,
    diseases_Id INTEGER NOT NULL,
    FOREIGN KEY (patient_Id) REFERENCES User(id),
    FOREIGN KEY (diseases_Id) REFERENCES diseases(id)
);

-- Insert default users
INSERT INTO User 
    (UserId, first_name, last_name,Gender, date_of_birth, email, hashed_password, Rule)
    VALUES 
    ('041206789', 'Mariam', 'Abbas','Female', '2004-12-15', 'random@example.com',
    '$2a$10$Jh3uNYj1l50.7eVFbmx6d.CqHZf/9TIL/z9D94qIIdgeAWY81lNEu', -- password is password123
    'admin');

    INSERT INTO User 
    (UserId, first_name, last_name,Gender, date_of_birth, email, hashed_password, Rule)
    VALUES 
    ('030309301', 'Fatima', 'Abbas','Female', '2003-03-27', 'random@example.com',
    '$2a$10$Jh3uNYj1l50.7eVFbmx6d.CqHZf/9TIL/z9D94qIIdgeAWY81lNEu', -- password is password123
    'patient');

    -- Insert default users
INSERT INTO Appointment 
    (Doctor_id,user_id,date_of_Appointment,TheTime)
    VALUES 
    ('1', '2', '2024-05-08','10:10');