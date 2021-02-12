package goal

import (
	"log"

	"github.com/skovsen/D2D_Protocol/workers"
	agentlogic "github.com/skovsen/D2D_AgentLogic"
	"github.com/paulmach/orb"
)

func EvaluteEndstate(goal agentlogic.Goal, position agentlogic.Vector) {
	log.Println("EVAL")
	log.Println(goal.Endgame)
	//implement different endgames here
	if goal.Endgame == "Gather" {
		goalPos := orb.Point{position.X, position.Y}
		goalLine := orb.MultiLineString{orb.LineString{goalPos}}
		workers.StateMux.Lock()
		//currPos := workers.MyState.Position

		m := agentlogic.Mission{
			Description: "Gather to goal",
			MissionType: workers.MyState.Mission.MissionType,
			Geometry:    goalLine,
		}
		workers.MyState.Mission = m
		workers.StateMux.Unlock()
		log.Println("Sending miss")
		workers.PrepareForMission(workers.MyState)
	}
}
