USE jumbotravel;

INSERT INTO flight_routes (airplanemapping_id, departure_country, arrival_country, departure_city, arrival_city, departure_airport, arrival_airport) VALUES
    (1, "ES", "ES", "PMI", "BCN", "PMI", "BCN"),
    (1, "ES", "ES", "BCN", "PMI", "BCN", "PMI"),
    (2, "ES", "ES", "MAD", "SVQ", "MAD", "SVQ"),
    (2, "ES", "ES", "SVQ", "MAD", "SVQ", "MAD"),
    (2, "ES", "ES", "MAD", "BCN", "MAD", "BCN"),
    (2, "ES", "ES", "BCN", "MAD", "BCN", "MAD"),
    (3, "ES", "ES", "MAD", "BIO", "MAD", "BIO"),
    (3, "ES", "ES", "BIO", "MAD", "BIO", "MAD"),
    (3, "ES", "ES", "MAD", "VLC", "MAD", "VLC"),
    (3, "ES", "ES", "VLC", "MAD", "VLC", "MAD");

INSERT INTO flights (route_id, status, departure_time, arrival_time) VALUES
    (1, "FLYING", TIMESTAMP("2022-03-20 10:05:00"), TIMESTAMP("2022-03-20 10:55:00")),
    (2, "FLYING", TIMESTAMP("2022-03-21 10:05:00"), TIMESTAMP("2022-03-21 10:55:00")),
    (3, "FLYING", TIMESTAMP("2022-03-22 10:05:00"), TIMESTAMP("2022-03-22 10:55:00")),
    (4, "FLYING", TIMESTAMP("2022-03-23 10:05:00"), TIMESTAMP("2022-03-23 10:55:00")),
    (5, "FLYING", TIMESTAMP("2022-03-24 10:05:00"), TIMESTAMP("2022-03-24 10:55:00")),
    (6, "FLYING", TIMESTAMP("2022-03-25 10:05:00"), TIMESTAMP("2022-03-25 10:55:00")),
    (7, "FLYING", TIMESTAMP("2022-03-26 10:05:00"), TIMESTAMP("2022-03-26 10:55:00")),
    (8, "FLYING", TIMESTAMP("2022-03-27 10:05:00"), TIMESTAMP("2022-03-27 10:55:00")),
    (9, "FLYING", TIMESTAMP("2022-03-28 10:05:00"), TIMESTAMP("2022-03-28 10:55:00"));

INSERT INTO flight_agents (flight_id, agent_id, agentmapping_id) VALUES
    (1, 1, 1),
    (2, 2, 2),
    (3, 3, 3),
    (4, 4, 4),
    (5, 5, 5),
    (6, 6, 6),
    (7, 7, 7),
    (8, 8, 8),
    (9, 9, 9);