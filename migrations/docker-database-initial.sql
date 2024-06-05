create table people(
    id serial primary key,
    name varchar,
    history varchar
);

INSERT INTO people(
	name, history)
	VALUES ('Afonso Pena', 'Afonso Augusto Moreira Pena foi um advogado e politico brasileiro. Foi o sexto presidente da República, de 1906 ate sua morte. Iniciou sua carreira politica durante o Imperio, exercendo varios cargos, incluindo de presidente de Minas Gerais, legislador, presidente do Banco da Republica e ministro de Estado.'),
	('Floriano Peixoto', 'Floriano Vieira Peixoto foi um militar e politico brasileiro, primeiro vice-presidente e segundo presidente do Brasil, cujo governo abrange a maior parte do periodo da historia brasileira conhecido como Republica da Espada');