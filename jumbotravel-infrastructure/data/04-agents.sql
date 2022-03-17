USE jumbotravel;

-- ASSISTANTS
INSERT INTO master_agents (dni, name, surname, email, password, type) VALUES
    ("111111111", "Pere", "Pons", "pere.pons@jumbotravel.com", "8ddae0e3780a6b9dcf84399fd092f7e2", "ASSISTANT"),
    ("222222222", "Catalina", "Tomás", "catalina.tomas@jumbotravel.com", "937354e62f9e2436c20f0c4bbbded9f4", "ASSISTANT"),
    ("333333333", "Margalida", "Clade", "margalida.clade@jumbotravel.com", "ec2824fec1a1e325ca2d29d9313244db", "ASSISTANT"),
    ("444444444", "Lola", "Herrera", "lola.herrera@jumbotravel.com", "4a01f9df732909dbb91401a043743b25", "ASSISTANT"),
    ("555555555", "Macarena", "Gómez", "macarena.gomez@jumbotravel.com", "959b08e08b7682ac3bb9865d466caa5e", "ASSISTANT"),
    ("666666666", "Carmen", "Pérez", "carmen.perez@jumbotravel.com", "a80f93c9aa4b01e65c0850698b01ffc9", "ASSISTANT"),
    ("777777777", "Iratxe", "Henzo", "iratxe.henzo@jumbotravel.com", "73bd4375932f105a59b9b5c748b616c3", "ASSISTANT"),
    ("888888888", "Jontxu", "Esquel", "jontxu.esquel@jumbotravel.com", "08e78d8b8266bc77fb40d098bcc14d15", "ASSISTANT"),
    ("999999999", "Amaia", "Petru", "amaia.petru@jumbotravel.com", "ce1a6df4e5c59c90dd4cb2ebc25b01e3", "ASSISTANT");

-- PROVIDERS
INSERT INTO master_agents (dni, name, surname, email, password, type, airport_id) VALUES
    ("121212121", "Oriol", "Piqué", "oriol.pique@jumbotravel.com", "f4c26ca217c15cc6a0ca5d06e675b97c", "PROVIDER", 2),
    ("131313131", "Manuel", "Becerra", "manuel.becerra@jumbotravel.com", "650cf3c5d47d90c3f17fb6c24c392b88", "PROVIDER", 4),
    ("141414141", "Juan", "Rosselló", "juan.rosello@jumbotravel.com", "5d677afa14d87aca2694654967051f0d", "PROVIDER", 1),
    ("151515151", "Bruno", "de Santiago", "bruno.santiago@jumbotravel.com", "c175924efff171a880efe46ec5266869", "PROVIDER", 3),
    ("161616161", "Koldo", "Arrebeitia", "koldo.arrebeitia@jumbotravel.com", "aabb1cb42e24b4ef695b77abe728fc6d", "PROVIDER", 6),
    ("171717171", "Xim", "Vergel", "xim.vergel@jumbotravel.com", "f2b993c9eb00f1b240ffe5d58753b556", "PROVIDER", 5);