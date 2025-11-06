package app

import (
	"fmt"

	"github.com/deikioveca/TheRedDevilsData/api/database"
	"github.com/deikioveca/TheRedDevilsData/api/football_client"
	"github.com/deikioveca/TheRedDevilsData/api/service"
	"github.com/spf13/cobra"
)


type CLI struct {
	Service service.Service
}


func NewCLI() *CLI {
	db 		:= database.InitDB()
	client 	:= football_client.NewFootballClient(db)
	service := service.NewService(db, client)

	return &CLI{
		Service: service,
	}
}


func (c *CLI) RootCmd() *cobra.Command {
	root := cobra.Command{
		Use: 	"football-cli",
		Short: 	"CLI for fetching and saving football data",
	}

	root.AddCommand(c.FetchCountriesCmd())
	root.AddCommand(c.FetchAllLeaguesForTeam())
	root.AddCommand(c.FetchTeam())
	root.AddCommand(c.FetchTeamStats())
	root.AddCommand(c.FetchVenues())
	root.AddCommand(c.FetchStandings())
	root.AddCommand(c.FetchFixtures())
	root.AddCommand(c.FetchInjuries())
	root.AddCommand(c.FetchSquad())

	return &root
}


func (c *CLI) FetchCountriesCmd() *cobra.Command {
	return &cobra.Command{
		Use: "fetch-countries",
		Short: "Fetch countries from API-Football and save them to DB",
		RunE: func(cmd *cobra.Command, args []string) error {
			countries, err := c.Service.SaveCountries()
			if err != nil {
				return err
			}
			for _, country := range countries.Response {
				fmt.Println("Saved:", country.Name, "(", country.Code, ")")
			}
			fmt.Println("All countries saved successfully!")
			return nil
		},
	}
}


func (c *CLI) FetchAllLeaguesForTeam() *cobra.Command {
	return &cobra.Command{
		Use: "fetch-all-leagues-for-team",
		Short: "Fetch all leagues in which Manchester United has played at least one match",
		RunE: func(cmd *cobra.Command, args []string) error {
			leagues, err := c.Service.SaveLeaguesForTeam()
			if err != nil {
				return err
			}
			fmt.Printf("Saved %d leagues for Manchester United\n", leagues.Results)
			return nil
		},
	}
}


func (c *CLI) FetchTeam() *cobra.Command {
	return &cobra.Command{
		Use: "fetch-team",
		Short: "Fetch manchester united from api-football",
		RunE: func(cmd *cobra.Command, args []string) error {
			team, err := c.Service.SaveTeam()
			if err != nil {
				return err
			}
			fmt.Printf("%+v", team)
			return nil
		},
	}
}


func (c *CLI) FetchTeamStats() *cobra.Command {
	return &cobra.Command{
		Use: "fetch-team-stats",
		Short: "Fetch and save manchester united stats for seasons: 2021, 2022, 2023",
		RunE: func(cmd *cobra.Command, args []string) error {
			stats, err := c.Service.SaveTeamStats()
			if err != nil {
				return err
			}

			for _, s := range stats {
				fmt.Printf("Saved stats for %s (%d) - Season %d\n", s.Response.Team.Name, s.Response.Team.ID, s.Response.League.Season)
			}

			fmt.Printf("Successfully saved %d team stats records.\n", len(stats))
			return nil
		},
	}
}


func (c *CLI) FetchVenues() *cobra.Command {
	return &cobra.Command{
		Use: "fetch-venues",
		Short: "Fetch and save all venues from england",
		RunE: func(cmd *cobra.Command, args []string) error {
			venues, err := c.Service.SaveVenues()
			if err != nil {
				return err
			}

			fmt.Printf("Successfully saved %d venues from england.\n", venues.Results)
			return nil
		},
	}
}


func (c *CLI) FetchStandings() *cobra.Command {
	return &cobra.Command{
		Use: "fetch-standings",
		Short: "Fetch and save manchester united standings for season 2021, 2022, 2023",
		RunE: func(cmd *cobra.Command, args []string) error {
			standings, err := c.Service.SaveStandings()
			if err != nil {
				return err
			}

			fmt.Printf("Successfully saved %d season standings for Manchester United.\n", len(standings))
			return nil
		},
	}
}


func (c *CLI) FetchFixtures() *cobra.Command {
	return &cobra.Command{
		Use: "fetch-fixtures",
		Short: "Fetch and save all manchester united fixtures in the premier league for 3 seasons 2021, 2022, 2023",
		RunE: func(cmd *cobra.Command, args []string) error {
			fixtures, err := c.Service.SaveFixtures()
			if err != nil {
				return err
			}

			fmt.Printf("Successfully saved %d season fixtures in the premier league for Manchester United.\n", len(fixtures))
			return nil
		},
	}
}


func (c *CLI) FetchInjuries() *cobra.Command {
	return &cobra.Command{
		Use: "fetch-injuries",
		Short: "Fetch and save all manchester united injuries for 3 seasons 2021, 2022, 2023",
		RunE: func(cmd *cobra.Command, args []string) error {
			injuries, err := c.Service.SaveInjuries()
			if err != nil {
				return err
			}

			fmt.Printf("Successfully saved %d season injuries for Manchester United.\n", len(injuries))
			return nil
		},
	}
}


func (c *CLI) FetchSquad() *cobra.Command {
	return &cobra.Command{
		Use: "fetch-squad",
		Short: "Fetch and save Manchester United squad for season 2025",
		RunE: func(cmd *cobra.Command, args []string) error {
			squad, err := c.Service.SaveSquad()
			if err != nil {
				return err
			}

			for _, dto := range squad.Response {
				for _, player := range dto.Players {
					fmt.Printf("%s: %s.\n", player.Name, player.Position)
				}
				fmt.Printf("Successfully saved %d Manchester United players.\n", len(dto.Players))
			}

			return nil
		},
	}
}