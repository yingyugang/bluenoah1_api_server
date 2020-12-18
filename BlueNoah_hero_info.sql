create table hero_info
(
    id        int auto_increment
        primary key,
    hero_name varchar(100)   not null,
    level     int default 1  null,
    user_id   int            not null,
    exp       int            null,
    atk       int default 10 not null,
    def       int default 1  not null,
    maxhp     int default 1  not null,
    hp        int default 3  not null,
    maxsp     int default 1  not null,
    sp        int default 1  not null,
    critical  int default 20 not null
);

INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (1, 'New hero', 1, 1, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (2, 'New hero', 1, 2, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (3, 'New hero', 1, 3, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (4, 'New hero', 1, 4, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (5, 'New hero', 1, 5, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (6, 'New hero', 1, 6, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (7, 'New hero', 1, 7, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (8, 'New hero', 1, 8, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (9, 'New hero', 1, 9, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (10, 'New hero', 1, 10, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (11, 'New hero', 1, 11, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (12, 'New hero', 1, 12, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (13, 'New hero', 1, 13, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (14, 'New hero', 1, 14, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (15, 'New hero', 1, 15, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (16, 'New hero', 1, 16, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (17, 'New hero', 1, 17, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (18, 'New hero', 1, 18, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (19, 'New hero', 1, 19, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (20, 'New hero', 1, 20, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (21, 'New hero', 1, 21, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (22, 'New hero', 1, 22, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (23, 'New hero', 1, 23, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (24, 'New hero', 1, 24, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (25, 'New hero', 1, 25, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (26, 'New hero', 1, 26, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (27, 'New hero', 1, 27, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (28, 'New hero', 1, 28, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (29, 'New hero', 1, 29, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (30, 'New hero', 1, 30, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (31, 'New hero', 1, 31, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (32, 'New hero', 1, 32, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (33, 'New hero', 1, 33, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (34, 'New hero', 1, 34, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (35, 'New hero', 1, 35, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (36, 'New hero', 1, 36, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (37, 'New hero', 1, 37, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (38, 'New hero', 1, 38, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (39, 'New hero', 1, 39, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (40, 'New hero', 1, 40, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (41, 'New hero', 1, 41, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (42, 'New hero', 1, 42, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (43, 'New hero', 1, 43, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (44, 'New hero', 1, 44, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (45, 'New hero', 1, 45, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (46, 'New hero', 1, 46, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (47, 'New hero', 1, 47, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (48, 'New hero', 1, 48, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (49, 'New hero', 1, 49, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (50, 'New hero', 1, 50, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (51, 'New hero', 1, 51, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (52, 'New hero', 1, 52, null, 10, 1, 1, 3, 1, 1, 20);
INSERT INTO BlueNoah.hero_info (id, hero_name, level, user_id, exp, atk, def, maxhp, hp, maxsp, sp, critical) VALUES (53, 'New hero', 1, 53, null, 10, 1, 1, 3, 1, 1, 20);