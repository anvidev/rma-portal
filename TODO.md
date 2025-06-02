# TODOs

### Road to v1.0-beta

 - [x] Add "Accept terms & conditions" checkbox to create form
 - [x] Add "I accept if the offer is above XX DKK" checkbox to create form
 - [x] Add Skancode logo to create page
 - [x] Add "wantQuote" data to ticket table
 - [x] Add new status "modtaget"
 - [x] Expand search field to billing information, model and serial number
 - [-] Ability to edit ticket on details page 
 - [x] Add "Vare er garanti", "Vare er ikke garanti", "Ved ikke" to create form
 - [x] Change "issue" min length to 20 characters
 - [x] Change filter for status to be multiple
 - [x] Fetch status and categories in frontend
 - [x] Show total in pagination
 - [x] Change ticket id to something like "00001-FKJSHF"
    - Autoincrement number combined with 5 random uppercase letters
 - [x] Update frontend for login page
 - [x] Modals for terms and coditions
 - [x] Ability to log out
 - [x] Show quote and warranty data on details page
 - [x] Thank you page after ticket creation (different redirect if user is logged in)

### Road to v1.1 (client facing)

 - [x] Ability for customers to look up a ticket / homepage
 - [x] Download PDF after ticket creation
 - [x] Email to skancode and customer when ticket is created
 - [x] Links in header to overview and create page
 - [ ] Add max height on ticket table
 - [ ] Add sorting/direction to filters
 - [x] Cache control for overview page
 - [x] Change redirect url for log out to login page
 - [x] Add a redirect searchparam to login page
 - [x] Fix caching after a ticket has been updated
 - [ ] Fix endpoints names with plural naming convention
 - [x] Fix pathname for ticket overview page to /admin/sager
 - [x] Rate limit middleware
 - [x] Abstract fetch call to api
