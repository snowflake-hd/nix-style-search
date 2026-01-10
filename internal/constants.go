package internal

const (
	// SEARCH_USERNAME | SEARCH_PASSWORD https://github.com/NixOS/nixos-search/blob/main/frontend/src/index.js
	// Just using what's already public ¯\_(ツ)_/¯
	SEARCH_USERNAME = "aWVSALXpZv"
	SEARCH_PASSWORD = "X8gPHnzL52wFEekuxsfQ9cSh"

	BASE_DOMAIN     = "https://search.nixos.org/backend"
	ALIAS_ENDPOINT  = BASE_DOMAIN + "/_cat/aliases"
	SEARCH_ENDPOINT = BASE_DOMAIN + "/%s/_search"

	INDEX_ERROR_TEMPLATE  = "Failed to load indices: %s\n\nEsc/Ctrl+C to quit.\n"
	INDEX_LOADING_MESSAGE = "Loading indices…\n\nEsc/Ctrl+C to quit.\n"
	INDEX_TITLE           = "Select an index\n\n"
	INDEX_HELP            = "\n↑/↓ move  •  Enter select  •  Esc/Ctrl+C quit\n"
	CURSOR_DISPLAY        = "  ▶ %s %s\n"

	SEARCH_ERROR_TEMPLATE      = "Error: %s\n\n"
	SEARCH_PROMPT              = "Type to search packages.\n"
	SEARCH_NO_RESULTS          = "No packages found.\n"
	SEARCH_RESULTS_TITLE       = "Packages:\n"
	SEARCH_DETAILS_TITLE       = "\nDetails:\n"
	SEARCH_DETAILS_PROMPT      = "\n(Press Tab to show details)\n"
	SEARCH_NAME_TEMPLATE       = "Name: %s\n"
	SEARCH_SUMMARY_TEMPLATE    = "Summary: %s\n"
	SEARCH_VERSION_TEMPLATE    = "Version: %s\n"
	SEARCH_ATTR_TEMPLATE       = "Attr: %s\n"
	SEARCH_HOMEPAGE_TEMPLATE   = "Homepage: %s\n"
	SEARCH_LICENSE_TEMPLATE    = "License: %s\n"
	SEARCH_NO_PACKAGE_SELECTED = "No package selected"
	SEARCH_SELECTION_OOB       = "Selection out of range"
	SEARCH_COPY_FAILED         = "Failed to copy to clipboard: %v\n"
	SEARCH_PACKAGE_SELECTED    = "Package %s selected\n"
	PAGE_SIZE                  = 10
)
