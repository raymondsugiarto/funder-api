INSERT INTO organization (id, name, code, origin, created_at, updated_at)
 VALUES ('1', '1', '1', 'localhost:9000', NOW(), NOW());    


INSERT INTO "user" (id, organization_id, user_type, created_at, updated_at)
 VALUES 
    ('1', '1', 'ADMIN', NOW(), NOW()),
    ('2', '1', 'OWNER', NOW(), NOW())
;    
-- admin dan owner

INSERT INTO user_credential (id, organization_id, user_id, username, password, created_at, updated_at)
 VALUES 
    ('1', '1', '1', 'admin', '$2a$07$rXsQxQHRwxwHYNTzHKTl.eilofdCZ9Ci0TTJmLdV6I7rxsYn/O74.', NOW(), NOW()), -- pass admin
    ('2', '1', '2', 'owner', '$2a$07$rXsQxQHRwxwHYNTzHKTl.eilofdCZ9Ci0TTJmLdV6I7rxsYn/O74.', NOW(), NOW()) -- pass admin
;
w