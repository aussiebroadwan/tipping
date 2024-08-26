CREATE TABLE competitions (
  id BIGINT PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

COMMENT ON COLUMN competitions.id IS 'Unique identifier for each competition';
COMMENT ON COLUMN competitions.name IS 'Name of the competition (e.g., NRL, NRLW)';

-- Insert initial competition records
INSERT INTO competitions (id, name) VALUES 
(111, 'NRL'),
(161, 'NRLW'),
(116, 'State of Origin'),
(156, 'Womens State of Origin');
