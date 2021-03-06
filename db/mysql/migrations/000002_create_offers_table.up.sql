CREATE TABLE IF NOT EXISTS Offers(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Title VARCHAR (255) ,
    Location VARCHAR (255) ,
    Description VARCHAR (10000) ,
    TitleImageURL VARCHAR (1000) ,
    UserID INT ,
    FOREIGN KEY (UserID) REFERENCES Users(ID) ,
    PRIMARY KEY (ID)
)