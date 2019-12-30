DROP TABLE IF EXISTS `mgram_cards`;

create table IF not exists `mgram_cards`(
    `id` int auto_increment,
    `facebook_id`   VARCHAR(1000) NOT NULL,
    `nickname`      VARCHAR(60) NOT NULL,
    `year`          INT(4),
    `month`         INT(2),
    `twitter_id`    VARCHAR(100),
    `team1`         VARCHAR(60),
    `team2`         VARCHAR(60),
    `word`          VARCHAR(60),
    `mgram1`        VARCHAR(60),
    `mgram2`        VARCHAR(60),
    `mgram3`        VARCHAR(60),
    `mgram4`        VARCHAR(60),
    `mgram5`        VARCHAR(60),
    `mgram6`        VARCHAR(60),
    `mgram7`        VARCHAR(60),
    `mgram8`        VARCHAR(60),
    `mgram9`        VARCHAR(60),
    `area`          VARCHAR(60),
    `card_color`    INT(1),
    PRIMARY KEY (`Id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

INSERT INTO mgram_cards(facebook_id, nickname, year, month, twitter_id, team1, team2, word, mgram1, mgram2, mgram3, mgram4, mgram5, mgram6, mgram7, mgram8, mgram9, area, card_color) VALUE ('testid','ramo4',2022,5,'ramo4','関西','エンジニア','一言一言','mgram1','mgram2','mgram3','mgram4','mgram5','mgram6','mgram7','mgram8','mgram9','osaka',1);