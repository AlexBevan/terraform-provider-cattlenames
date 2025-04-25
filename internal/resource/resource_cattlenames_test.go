package resource

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCattleNames_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { /* TODO: add if needed */ },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCattleNamesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCattleNamesBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCattleNamesExists("cattlenames.test"),
					// Verify name is set and matches the cattle names pattern
					resource.TestMatchResourceAttr(
						"cattlenames.test", "name",
						regexp.MustCompile(`^(Bessie|MooMoo|Daisy|Buttercup|Clarabelle|Elsie|Gertie|Mabel|Nellie|Rosie)$`),
					),
					// Verify family is set to default
					resource.TestCheckResourceAttr(
						"cattlenames.test", "family", "cattle",
					),
				),
			},
		},
	})
}

func TestAccCattleNames_got(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { /* TODO: add if needed */ },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCattleNamesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCattleNamesGOT,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCattleNamesExists("cattlenames.got_test"),
					// Verify family is set to got
					resource.TestCheckResourceAttr(
						"cattlenames.got_test", "family", "got",
					),
					// Verify name is a GOT character
					resource.TestMatchResourceAttr(
						"cattlenames.got_test", "name",
						regexp.MustCompile(`^(Jon Snow|Daenerys Targaryen|Tyrion Lannister|Arya Stark|Cersei Lannister|Jaime Lannister|Sansa Stark|Bran Stark|Jorah Mormont|Theon Greyjoy|Night King)$`),
					),
				),
			},
		},
	})
}

func TestAccCattleNames_27club(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { /* TODO: add if needed */ },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCattleNamesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCattleNames27Club,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCattleNamesExists("cattlenames.club_test"),
					// Verify family is set to 27club
					resource.TestCheckResourceAttr(
						"cattlenames.club_test", "family", "27club",
					),
					// Verify name is a 27 Club member
					resource.TestMatchResourceAttr(
						"cattlenames.club_test", "name",
						regexp.MustCompile(`^(Kurt Cobain|Amy Winehouse|Jimi Hendrix|Janis Joplin|Jim Morrison|Robert Johnson|Brian Jones|Eazy-E)$`),
					),
				),
			},
		},
	})
}

func TestAccCattleNames_caseInsensitive(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { /* TODO: add if needed */ },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCattleNamesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCattleNamesUpperCase,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCattleNamesExists("cattlenames.case_test"),
					// Verify name is a GOT character despite uppercase family
					resource.TestMatchResourceAttr(
						"cattlenames.case_test", "name",
						regexp.MustCompile(`^(Jon Snow|Daenerys Targaryen|Tyrion Lannister|Arya Stark|Cersei Lannister|Jaime Lannister|Sansa Stark|Bran Stark|Jorah Mormont|Theon Greyjoy|Night King)$`),
					),
				),
			},
		},
	})
}

func testAccCheckCattleNamesExists(n string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[n]
        if !ok {
            return fmt.Errorf("Resource not found: %s", n)
        }
        if rs.Primary.ID == "" {
            return fmt.Errorf("Resource ID not set")
        }
        
        // Verify required attributes are set
        if name := rs.Primary.Attributes["name"]; name == "" {
            return fmt.Errorf("name attribute not set")
        }
        if family := rs.Primary.Attributes["family"]; family == "" {
            return fmt.Errorf("family attribute not set")
        }
        return nil
    }
}

func testAccCheckCattleNamesDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cattlenames" {
			continue
		}
		// In a real provider, we would verify the resource is actually destroyed
		// In our case, since it's just generating names, there's nothing to destroy
	}
	return nil
}

const testAccCheckCattleNamesBasic = ` 
resource "cattlenames" "test" {
}
`

const testAccCheckCattleNamesGOT = `
resource "cattlenames" "got_test" {
  family = "got"
}
`

const testAccCheckCattleNames27Club = `
resource "cattlenames" "club_test" {
  family = "27club"
}
`

const testAccCheckCattleNamesUpperCase = `
resource "cattlenames" "case_test" {
  family = "GOT"
}
`

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
    testAccProvider = &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{
            "cattlenames": ResourceCattleNames(),
        },
    }
    testAccProviders = map[string]*schema.Provider{
        "cattlenames": testAccProvider,
    }
}
