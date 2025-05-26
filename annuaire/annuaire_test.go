package annuaire

import (
	"fmt"
	"testing"
)

func TestAjouterContact(t *testing.T) {
	type testCase struct {
		nom         string
		tel         string
		expectedMsg string
		expectedErr error
	}

	tests := map[string]testCase{
		"Valid contact": {
			nom:         "Alice",
			tel:         "0811223344",
			expectedMsg: "Contact Alice ajouté avec succès",
			expectedErr: nil,
		},
		"Empty name": {
			nom:         "",
			tel:         "0811223344",
			expectedMsg: "",
			expectedErr: fmt.Errorf("nom et numéro de téléphone requis"),
		},
		"Empty phone": {
			nom:         "Alice",
			tel:         "",
			expectedMsg: "",
			expectedErr: fmt.Errorf("nom et numéro de téléphone requis"),
		},
		"Duplicate contact": {
			nom:         "Charlie",
			tel:         "0811223344",
			expectedMsg: "",
			expectedErr: fmt.Errorf("le contact Charlie existe déjà"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			msg, err := AjouterContact(tc.nom, tc.tel)
			if tc.expectedErr != nil {
				if err == nil || err.Error() != tc.expectedErr.Error() {
					t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
				}
				if msg != "" {
					t.Errorf("Expected empty message, got %s", msg)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if msg != tc.expectedMsg {
					t.Errorf("Expected message %s, got %s", tc.expectedMsg, msg)
				}
			}
		})
	}
	ResetContacts()
}

func TestRechercherContact(t *testing.T) {
	type testCase struct {
		nom         string
		expectedMsg string
		expectedErr error
	}

	tests := map[string]testCase{
		"Valid contact": {
			nom:         "Charlie",
			expectedMsg: "Contact Charlie trouvé : 0811223344",
			expectedErr: nil,
		},
		"Non-existent contact": {
			nom:         "NonExistent",
			expectedMsg: "",
			expectedErr: fmt.Errorf("contact NonExistent non trouvé"),
		},
		"Empty name": {
			nom:         "",
			expectedMsg: "",
			expectedErr: fmt.Errorf("nom requis pour rechercher un contact"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			msg, err := RechercherContact(tc.nom)
			if tc.expectedErr != nil {
				if err == nil || err.Error() != tc.expectedErr.Error() {
					t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
				}
				if msg != "" {
					t.Errorf("Expected empty message, got %s", msg)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if msg != tc.expectedMsg {
					t.Errorf("Expected message %s, got %s", tc.expectedMsg, msg)
				}
			}
		})
	}
	ResetContacts()
}

func TestListerContacts(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "List contacts",
			expected: "Liste des contacts :\nCharlie : 0811223344\nDave : 0811333445\nEve : 0811445566\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ListerContacts()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
	ResetContacts()
}

func TestListerContactsEmpty(t *testing.T) {
	// Reset contacts to an empty state for this test
	EmptyContacts()
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "No contacts",
			expected: "Aucun contact trouvé",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ListerContacts()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
	// Reset contacts back to initial state
	ResetContacts()
}

func TestMettreAJourContact(t *testing.T) {
	type testCase struct {
		nom         string
		tel         string
		expectedMsg string
		expectedErr error
	}

	tests := map[string]testCase{
		"Valid update": {
			nom:         "Charlie",
			tel:         "0711223344",
			expectedMsg: "Contact Charlie mis à jour avec succès",
			expectedErr: nil,
		},
		"Non-existent contact": {
			nom:         "NonExistent",
			tel:         "0811223344",
			expectedMsg: "",
			expectedErr: fmt.Errorf("le contact NonExistent n'existe pas"),
		},
		"Empty name": {
			nom:         "",
			tel:         "0811223344",
			expectedMsg: "",
			expectedErr: fmt.Errorf("nom et numéro de téléphone requis pour mettre à jour un contact"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			msg, err := MettreAJourContact(tc.nom, tc.tel)
			if tc.expectedErr != nil {
				if err == nil || err.Error() != tc.expectedErr.Error() {
					t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
				}
				if msg != "" {
					t.Errorf("Expected empty message, got %s", msg)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if msg != tc.expectedMsg {
					t.Errorf("Expected message %s, got %s", tc.expectedMsg, msg)
				}
			}
		})
	}
	ResetContacts()
}

func TestSupprimerContact(t *testing.T) {
	type testCase struct {
		nom         string
		expectedMsg string
		expectedErr error
	}

	tests := map[string]testCase{
		"Valid deletion": {
			nom:         "Charlie",
			expectedMsg: "Contact Charlie supprimé avec succès",
			expectedErr: nil,
		},
		"Non-existent contact": {
			nom:         "NonExistent",
			expectedMsg: "",
			expectedErr: fmt.Errorf("le contact NonExistent n'existe pas"),
		},
		"Empty name": {
			nom:         "",
			expectedMsg: "",
			expectedErr: fmt.Errorf("nom requis pour supprimer un contact"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			msg, err := SupprimerContact(tc.nom)
			if tc.expectedErr != nil {
				if err == nil || err.Error() != tc.expectedErr.Error() {
					t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
				}
				if msg != "" {
					t.Errorf("Expected empty message, got %s", msg)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if msg != tc.expectedMsg {
					t.Errorf("Expected message %s, got %s", tc.expectedMsg, msg)
				}
			}
		})
	}
	ResetContacts()
}
