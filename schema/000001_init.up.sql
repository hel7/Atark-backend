CREATE TABLE User (
                      UserID INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                      Username VARCHAR(255) UNIQUE NOT NULL,
                      Email VARCHAR(255) UNIQUE NOT NULL,
                      Password VARCHAR(255) NOT NULL,
                      Role ENUM('User','Admin') NOT NULL
);

CREATE TABLE Feed (
                      FeedID INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                      FeedName VARCHAR(255) UNIQUE NOT NULL,
                      Quantity INT NOT NULL
);

CREATE TABLE Animal (
                        AnimalID INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                        AnimalName VARCHAR(255) NOT NULL,
                        Number INT NOT NULL UNIQUE,
                        DateOfBirth DATE NOT NULL,
                        Sex ENUM('M', 'F') NOT NULL,
                        Age INT NOT NULL,
                        MedicalInfo TEXT NOT NULL
);

CREATE TABLE Farm (
                      FarmID INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                      UserID INT UNSIGNED NOT NULL,
                      FarmName VARCHAR(255) UNIQUE NOT NULL,
                      FOREIGN KEY (UserID) REFERENCES User(UserID) ON DELETE CASCADE
);

CREATE TABLE Activity (
                          ActivityID INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                          AnimalID INT UNSIGNED,
                          ActivityType VARCHAR(255) NOT NULL,
                          StartTime DATETIME NOT NULL,
                          EndTime DATETIME NOT NULL,
                          Latitude DOUBLE NOT NULL,
                          Longitude DOUBLE NOT NULL,
                          FOREIGN KEY (AnimalID) REFERENCES Animal(AnimalID) ON DELETE CASCADE
);

CREATE TABLE Biometrics (
                            BiometricID INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                            AnimalID INT UNSIGNED,
                            Pulse INT NOT NULL,
                            Temperature DOUBLE NOT NULL,
                            Weight DOUBLE NOT NULL,
                            BreathingRate INT NOT NULL,
                            FOREIGN KEY (AnimalID) REFERENCES Animal(AnimalID) ON DELETE CASCADE
);

CREATE TABLE FeedingSchedule (
                                 ScheduleID INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                                 AnimalID INT UNSIGNED,
                                 FeedID INT UNSIGNED,
                                 FeedingTime TIME NOT NULL,
                                 FeedingDate DATE NOT NULL,
                                 AllocatedQuantity INT NOT NULL,
                                 FOREIGN KEY (AnimalID) REFERENCES Animal(AnimalID) ON DELETE CASCADE,
                                 FOREIGN KEY (FeedID) REFERENCES Feed(FeedID) ON DELETE CASCADE
);


CREATE TABLE FarmAnimal (
                            FarmID INT UNSIGNED NOT NULL,
                            AnimalID INT UNSIGNED NOT NULL,
                            PRIMARY KEY (FarmID, AnimalID),
                            FOREIGN KEY (FarmID) REFERENCES Farm(FarmID) ON DELETE CASCADE,
                            FOREIGN KEY (AnimalID) REFERENCES Animal(AnimalID) ON DELETE CASCADE
);