package resource

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCattleNames() *schema.Resource {
	return &schema.Resource{
		Create: resourceCattleNamesCreate,
		Read:   resourceCattleNamesRead,
		Delete: resourceCattleNamesDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true, // Ensure name is generated and not user-settable
			},
			"family": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cattle", // Default family is 'cattle'
				ForceNew: true, // Changing family forces a new resource
			},
		},
	}
}

func resourceCattleNamesCreate(d *schema.ResourceData, m interface{}) error {
	family := d.Get("family").(string) // Retrieve the family parameter
	cattleName := generateCattleName(family)
	d.Set("name", cattleName) // Ensure the name field is set to the generated random name
	d.SetId(fmt.Sprintf("%d", time.Now().UnixNano()))
	return nil
}

func resourceCattleNamesRead(d *schema.ResourceData, m interface{}) error {
	if d.Id() == "" {
		d.SetId("")
		return nil
	}
	return nil
}

func resourceCattleNamesDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func generateCattleName(family string) string {
	var names []string

	family = strings.ToLower(family) // Make family case-insensitive

	switch family {
	case "got":
		names = []string{
			"Jon Snow", "Daenerys Targaryen", "Tyrion Lannister", "Arya Stark", "Cersei Lannister",
			"Jaime Lannister", "Sansa Stark", "Bran Stark", "Jorah Mormont", "Theon Greyjoy",
			"Eddard Stark", "Robert Baratheon", "Stannis Baratheon", "Renly Baratheon", "Melisandre",
			"Samwell Tarly", "Gilly", "Brienne of Tarth", "Sandor Clegane", "Gregor Clegane",
			"Petyr Baelish", "Varys", "Shae", "Ygritte", "Tormund Giantsbane",
			"Margaery Tyrell", "Loras Tyrell", "Olenna Tyrell", "Joffrey Baratheon", "Tommen Baratheon",
			"Myrcella Baratheon", "Roose Bolton", "Ramsay Bolton", "Balon Greyjoy", "Euron Greyjoy",
			"Daario Naharis", "Missandei", "Grey Worm", "Khal Drogo", "Viserys Targaryen",
			"Aerys II Targaryen", "Rhaegar Targaryen", "Lyanna Stark", "Benjen Stark", "Hodor",
			"Gendry", "Podrick Payne", "Jaqen H'ghar", "Beric Dondarrion", "Thoros of Myr",
			"Ellaria Sand", "Oberyn Martell", "Doran Martell", "Trystane Martell", "Hot Pie",
			"Maester Aemon", "Jeor Mormont", "Alliser Thorne", "Eddison Tollett", "Yara Greyjoy",
			"Meera Reed", "Jojen Reed", "Brynden Tully", "Edmure Tully", "Walder Frey",
			"Lysa Arryn", "Robin Arryn", "Shireen Baratheon", "Selyse Baratheon", "Maester Pycelle",
			"Qyburn", "Kevan Lannister", "Brynden Rivers", "Leaf", "Night King",
		}
	case "27club":
		names = []string{
			"Kurt Cobain", "Amy Winehouse", "Jimi Hendrix", "Janis Joplin", "Jim Morrison",
			"Robert Johnson", "Brian Jones", "Ron \"Pigpen\" McKernan", "Kristen Pfaff", "Pete Ham",
			"Chris Bell", "D. Boon", "Jean-Michel Basquiat", "Richey Edwards", "Mia Zapata",
			"Alan Wilson", "Dave Alexander", "Jeremy Michael Ward", "Les Harvey", "Alexander Bashlachev",
			"Arlester \"Dyke\" Christian", "Jesse Belvin", "Linda Jones", "Malcolm Hale", "Nat Jaffe",
			"Rudy Lewis", "Stuart Sutcliffe", "Jacob Miller", "Gary Thain", "Helmut Köllen",
			"Kristen Pfaff", "Sean McCabe", "Freaky Tah", "Fat Pat", "Stretch",
			"The Notorious B.I.G.", "Tupac Shakur", "Big L", "Big Pun", "Eazy-E",
		}
	default: // Default to 'cattle'
		names = []string{
			"Bessie", "MooMoo", "Daisy", "Buttercup", "Clarabelle",
			"Elsie", "Gertie", "Mabel", "Nellie", "Rosie",
		}
	}

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return names[randGen.Intn(len(names))]
}