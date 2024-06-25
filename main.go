package main

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "fmt"
    "io"
    "os"
    "strconv"
    "encoding/xml"
    "github.com/BenLubar/df2014/cp437"
)

type World struct {
    // XMLName xml.Name `xml:"df_world"`
    Regions []struct {
        ID      int32        `xml:"id"`
        Name    string      `xml:"name"`
        Type    string      `xml:"type"`
    } `xml:"regions>region"`
    UndergroundRegions []struct {
        ID      int32        `xml:"id"`
        Type    string      `xml:"type"`
        Depth   int32       `xml:"depth"`
    } `xml:"underground_regions>underground_region"`
    Sites []struct {
        ID          int32        `xml:"id"`
        Name        string      `xml:"name"`
        Type        string      `xml:"type"`
        Cords       string      `xml:"cords"`
        Rectangle   string      `xml:"rectangle"`
        Structures  struct {
            ID      int32        `xml:"local_id"`
            Type    string      `xml:"type"`
            Name    string      `xml:"name"`
        } `xml:"structures"`
    } `xml:"sites>site"`
    Artifacts []struct {
        ID          int32        `xml:"id"`
        Name        string      `xml:"name"`
        Item        struct {
            Name    string      `xml:"name_string"`
        } `xml:"item"`
        SiteID      int32        `xml:"site_id"`
    } `xml:"artifacts>artifact"`
    HistoricalFigures   []struct {
        ID          int32            `xml:"id"`
        Name        string          `xml:"name"`
        Race        string          `xml:"race"`
        Gender      string          `xml:"caste"`
        Appeared    int32           `xml:"appeared"`
        Birth       int32           `xml:"birth_year"`
        Death       int32           `xml:"death_year"`
        Job         string          `xml:"associated_type"`
        Spheres     []string        `xml:"sphere"`
        Family      []struct {
            Type    string      `xml:"link_type"`
            ID      int32        `xml:"hfid"`
        } `xml:"hf_link"`
        EntityRelation []struct {
            Type    string      `xml:"link_type"`
            ID      int32        `xml:"entity_id"`
            Strength int32      `xml:"link_strength"`
        } `xml:"entity_link"`
        RelationPos []struct {
            PositionProfileID   int32        `xml:"position_profile_id"`
            EntityID            int32        `xml:"entity_id"`
            StartYear           int32        `xml:"start_year"`
            EndYear             int32        `xml:"end_year"`
        }
        Skills      []struct {
            Name    string      `xml:"skill"`
            Total   int32       `xml:"total_ip"`
        } `xml:"hf_skill"`
        Goal        string          `xml:"goal"`
        Disguise    int32            `xml:"used_identity_id"`
        Knowledge   []string        `xml:"intercation_knowledge"`
        Plots       []struct {
            ID      int32        `xml:"local_id"`
            Type    string      `xml:"type"`
        } `xml:"intrigue_plot"`
        Recidence   int32            `xml:"ent_pop_id"`
    } `xml:"historical_figures>historical_figure"`
    EntityPopulations   []struct {
        ID          int32        `xml:"id"`
    } `xml:"entity_populations>entity_population"`
    Entities        []struct {
        ID          int32        `xml:"id"`
        Name        string      `xml:"name"`
    } `xml:"entities>entity"`
    HistoricalEvents []HistoricalEvent `xml:"historical_events>historical_event"`
    HistoricalEventCollections []struct {
        ID          int32   `xml:"id"`
        Start       int32   `xml:"start_year"`
        StartTime   int32   `xml:"start_seconds72"`
        End         int32   `xml:"end_year"`
        EndTime     int32   `xml:"end_seconds72"`
        Events      []int32     `xml:"event"`
        Type        string  `xml:"type"`
        ParentEvent int32   `xml:"parent_eventcol"`
        Times       int32   `xml:"ordinal"`
        Site        int32   `xml:"site_id"`
        Coords       string  `xml:"coords"`
        Attacking   int32   `xml:"attacking_enid"`
        Defending   int32   `xml:"defending_enid"`
        Name        string  `xml:"name"`
        WarEvent    int32   `xml:"war_eventcol"`
        Attackers   []int32     `xml:"attacking_hfid"`
        Defenders   []int32     `xml:"defending_hfid"`
        AttackingSquadsRace    []string  `xml:"attacking_squad_race"`
        AttackingSquadsFrom    []int32   `xml:"attacking_squad_entity_pop"`
        AttackingSquadsNumber  []int32   `xml:"attacking_squad_number"`
        AttackingSquadsDeaths  []int32   `xml:"attacking_squad_deaths"`
        AttackingSquadsSite    []int32   `xml:"attacking_squad_site"`
        DefendingSquadsRace    []string  `xml:"defending_squad_race"`
        DefendingSquadsFrom    []int32   `xml:"defending_squad_entity_pop"`
        DefendingSquadsNumber  []int32   `xml:"defending_squad_number"`
        DefendingSquadsDeaths  []int32   `xml:"defending_squad_deaths"`
        DefendingSquadsSite    []int32   `xml:"defending_squad_site"`
        Outcome     string  `xml:"outcome"`
        Agressor    int32   `xml:"aggressor_ent_id"`
        Defensive   int32   `xml:"defender_ent_id"`
    } `xml:"historical_event_collections>historical_event_collection"`
    HistoricalEras []struct {
        Name        string  `xml:"name"`
        Start       int32    `xml:"start_year"`
    } `xml:"historical_eras>historical_era"`
    WrittenContents []struct {
        ID          int32   `xml:"id"`
        Title       string  `xml:"title"`
        Author      int32   `xml:"author_hfid"`
        Roll        int32   `xml:"author_roll"`
        Form        string  `xml:"form"`
        FormID      int32    `xml:"form_id"`
        Style       []string  `xml:"style"`
    } `xml:"written_contents>written_content"`
    PoeticForms []struct {
        ID          int32    `xml:"id"`
        Description string  `xml:"description"`
    } `xml:"poetic_forms>poetic_form"`
    MusicalForms []struct {
        ID          int32    `xml:"id"`
        Description string  `xml:"description"`
    } `xml:"musical_forms>musical_form"`
    DanceForms []struct {
        ID          int32    `xml:"id"`
        Description string  `xml:"description"`
    } `xml:"dance_forms>dance_form"`
}

type HistoricalEvent struct {
    ID          int32        `xml:"id"`
    Year        int32        `xml:"year"`
    Time        int32       `xml:"seconds72"`
    Type        string      `xml:"type"`
    Civ         int32       `xml:"civ_id"`
    Link        string      `xml:"link"`
    HFID        int32       `xml:"hfid"`
    Position    int32       `xml:"position_id"`
    HFIDTarget  int32       `xml:"hfid_target"`
    Artifact    int32       `xml:"artifact_id"`
    Unit        int32   `xml:"unit_id"`
    Creator     int32   `xml:"hist_figure_id"`
    Attacker    int32   `xml:"attacker_hfid"`
    AttackerCiv int32   `xml:"attacker_civ_id"`
    DefenderCiv int32   `xml:"defender_civ_id"`
    SiteCiv     int32   `xml:"site_civ_id"`
    Site        int32  `xml:"site_id"`
    AttackerGeneral int32   `xml:"attacker_general_hfid"`
    DefenderGeneral int32   `xml:"defender_general_hfid"`
    Coords      string  `xml:"coords"`
    State       string  `xml:"state"`
    Changee     int32   `xml:"changee_hfid"`
    Changer     int32   `xml:"changer_hfid"`
    OldRace     string  `xml:"old_race"`
    OldGender   string  `xml:"old_caste"`
    NewRace     string  `xml:"new_race"`
    NewGender   string  `xml:"new_caste"`
    Structure   int32   `xml:"structure_id"`
    WCID        int32   `xml:"wcid"`
    Master      int32   `xml:"master_wcid"`
    Site1       int32   `xml:"site_id1"`
    Site2       int32   `xml:"site_id2"`
    Entity      int32   `xml:"entity_id"`
    SnatchedTarget int32    `xml:"target_hfid"`
    Snatcher    int32   `xml:"snatcher_hfid"`
    Slayer      int32   `xml:"slayer_hfid"`
    SlayerRace  string  `xml:"slayer_race"`
    SlayerGender string `xml:"slayer_caste"`
    SlayerItem  int32   `xml:"slayer_item_id"`
    SlayerGun   int32   `xml:"slayer_shooter_item_id"`
    Cause       string  `xml:"cause"`
    Group       int32   `xml:"group_hfid"`
    Group1      int32   `xml:"group_1_hfid"`
    Group2      int32   `xml:"group_2_hfid"`
    Actor       int32   `xml:"actor_hfid"`
    Subtype     string  `xml:"subtype"`
    Woundee     int32   `xml:"woundee_hfid"`
    Wounder     int32   `xml:"wounder_hfid"`
    NewSite     int32   `xml:"new_site_civ_id"`
    NewLeader   int32   `xml:"new_leader_hfid"`
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var world World

// postAlbums adds an album from JSON received in the request body.
// func postAlbums(c *gin.Context) {
//     var newAlbum album

//     // Call BindJSON to bind the received JSON to
//     // newAlbum.
//     if err := c.BindJSON(&newAlbum); err != nil {
//         return
//     }

//     // Add the new album to the slice.
//     albums = append(albums, newAlbum)
//     c.IndentedJSON(http.StatusCreated, newAlbum)
// }

func getAll(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world)
}

func getRegions(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.Regions)
}

func getUndergroundRegions(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.UndergroundRegions)
}

func getSites(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.Sites)
}

func getArtifacts(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.Artifacts)
}

func getHistoricalFigures(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.HistoricalFigures)
}

func getEntityPopulations(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.EntityPopulations)
}

func getEntities(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.Entities)
}

func getHistoricalEvents(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.HistoricalEvents)
}

func getHistoricalEventCollections(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.HistoricalEventCollections)
}

func getHistoricalEras(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.HistoricalEras)
}

func getWrittenContents(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.WrittenContents)
}

func getPoeticForms(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.PoeticForms)
}

func getMusicalForms(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.MusicalForms)
}

func getDanceForms(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, world.DanceForms)
}

func getRegionByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.Regions {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "region not found"})
}

func getUndergroundRegionByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.UndergroundRegions {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "underground region not found"})
}

func getSiteByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.Sites {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "region not found"})
}

func getArtifactByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.Artifacts {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "region not found"})
}

func getHistoricalFigureByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.HistoricalFigures {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "region not found"})
}

func getEntityByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.Entities {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "region not found"})
}

func getHistoricalEventByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.HistoricalEvents {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "region not found"})
}

func getHistoricalEventCollectionByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.HistoricalEventCollections {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "region not found"})
}

func getWrittenContentByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.WrittenContents {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "region not found"})
}

func getPoeticFormByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.PoeticForms {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "region not found"})
}

func getMusicalFormByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.MusicalForms {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "region not found"})
}

func getDanceFormByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range world.DanceForms {
        if strconv.Itoa(int(a.ID)) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "region not found"})
}

func getRegionSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.Regions))
}

func getUndergroundRegionSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.UndergroundRegions))
}

func getSiteSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.Sites))
}

func getArtifactSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.Artifacts))
}

func getHistoricalFigureSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.HistoricalFigures))
}

func getEntityPopulationSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.EntityPopulations))
}

func getEntitySize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.Entities))
}

func getHistoricalEventSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.HistoricalEvents))
}

func getHistoricalEventCollectionSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.HistoricalEventCollections))
}

func getHistoricalEraSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.HistoricalEras))
}

func getWrittenContentSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.WrittenContents))
}

func getPoeticFormSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.PoeticForms))
}

func getMusicalFormSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.MusicalForms))
}

func getDanceFormSize(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, len(world.DanceForms))
}

func getRegionNames(c *gin.Context) {
    li := []string{}

    for _, a := range world.Regions {
        if a.Name != "" {
            li = append(li, a.Name)
        }
    }

    c.IndentedJSON(http.StatusOK, li)
}

func getEntityEvents(c *gin.Context) {
    id := c.Param("id")
    li := []HistoricalEvent{}

    for _, a := range world.HistoricalEvents {
        if a.Entity == id {
            li = append(li, a)
        }
    }

    c.IndentedJSON(http.StatusOK, li)
}

func getSiteEvents(c *gin.Context) {
    id := c.Param("id")
    li := []HistoricalEvent{}

    for _, a := range world.HistoricalEvents {
        if a.Site == id {
            li = append(li, a)
        }
    }

    c.IndentedJSON(http.StatusOK, li)
}

func getHistoricalFigureEvents(c *gin.Context) {
    id := c.Param("id")
    li := []HistoricalEvent{}

    for _, a := range world.HistoricalEvents {
        if a.HFID == id || a.Creator || a.Changee ||
            a.Changer || a.HFIDTarget || a.SnatchedTarget ||
            a.Snatcher || a.Slayer || a.Woundee || a.Wounder {
            li = append(li, a)
        }
    }

    c.IndentedJSON(http.StatusOK, li)
}

func identReader(encoding string, input io.Reader) (io.Reader, error) {
    return input, nil
}

func main() {
    xmlFile, err := os.Open("legends.xml")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Succesfully opened file")
    defer xmlFile.Close()

    r := cp437.NewReader(xmlFile)
    d := xml.NewDecoder(r)
    // empty decoder due to the lake of CP437 info
    d.CharsetReader = identReader
    err = d.Decode(&world)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Succesfully read file")

	router := gin.Default()
	router.GET("/all", getAll)
    router.GET("/regions", getRegions)
    router.GET("/underground_regions", getUndergroundRegions)
    router.GET("/sites", getSites)
    router.GET("/artifacts", getArtifacts)
    router.GET("/historical_figures", getHistoricalFigures)
    router.GET("/entity_populations", getEntityPopulations)
    router.GET("/entities", getEntities)
    router.GET("/historical_events", getHistoricalEvents)
    router.GET("/historical_event_collections", getHistoricalEventCollections)
    router.GET("/historical_eras", getHistoricalEras)
    router.GET("/written_contents", getWrittenContents)
    router.GET("/poetic_forms", getPoeticForms)
    router.GET("/musical_forms", getMusicalForms)
    router.GET("/dance_forms", getDanceForms)
    router.GET("/region/names", getRegionNames)
    router.GET("/region/:id", getRegionByID)
    router.GET("/underground_region/:id", getUndergroundRegionByID)
    router.GET("/site/:id", getSiteByID)
    router.GET("/artifact/:id", getArtifactByID)
    router.GET("/historical_figure/:id", getHistoricalFigureByID)
    router.GET("/entity/:id", getEntityByID)
    router.GET("/historical_event/:id", getHistoricalEventByID)
    router.GET("/historical_event_collection/:id", getHistoricalEventCollectionByID)
    router.GET("/written_content/:id", getWrittenContentByID)
    router.GET("/poetic_form/:id", getPoeticFormByID)
    router.GET("/musical_form/:id", getMusicalFormByID)
    router.GET("/dance_form/:id", getDanceFormByID)
    router.GET("/size/regions", getRegionSize)
    router.GET("/size/underground_regions", getUndergroundRegionSize)
    router.GET("/size/sites", getSiteSize)
    router.GET("/size/artifacts", getArtifactSize)
    router.GET("/size/historical_figures", getHistoricalFigureSize)
    router.GET("/size/entity_populations", getEntityPopulationSize)
    router.GET("/size/entities", getEntitySize)
    router.GET("/size/historical_events", getHistoricalEventSize)
    router.GET("/size/historical_event_collections", getHistoricalEventCollectionSize)
    router.GET("/size/historical_eras", getHistoricalEraSize)
    router.GET("/size/written_contents", getWrittenContentSize)
    router.GET("/size/poetic_forms", getPoeticFormSize)
    router.GET("/size/musical_forms", getMusicalFormSize)
    router.GET("/size/dance_forms", getDanceFormSize)

    router.Run("localhost:8080")
}
