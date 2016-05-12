// graphjson tests - just make sure the json conversions work
package graphneo

import (
	"github.com/adrianco/spigo/tooling/archaius"
	"github.com/adrianco/spigo/tooling/names"
	"testing"
	"time"
)

// reader parses graphjson
func TestGraph(t *testing.T) {
	testNeo := `
	  CREATE (test_mysql00:mysql {instance:"mysql00", name:"test.us-east-1.zoneA..mysql00...mysql.store", package:"store", timestamp:"2016-04-17T13:40:05.938437713-07:00", ip:"54.198.0.1", region:"us-east-1", zone: "zoneA"}),
          (test_mysql01:mysql {instance:"mysql01", name:"test.us-east-1.zoneA..mysql01...mysql.store", package:"store", timestamp:"2016-04-17T13:40:05.938513762-07:00", ip:"54.221.0.1", region:"us-east-1", zone: "zoneA"}),
          (test_mysql00)-[:CONNECTION]->(test_mysql01)
                `
	archaius.Conf.Arch = "test"
	Setup(archaius.Conf.Arch)
	Write(testNeo)
	dal0 := names.Make("test", "us-east-1", "ZoneA", "dal", "staash", 0)
	WriteNode(dal0+" staash", time.Now())
	WriteEdge(dal0+" test.us-east-1.zoneA..mysql00...mysql.store", time.Now())
	WriteEdge(dal0+" test.us-east-1.zoneA..mysql01...mysql.store", time.Now())
	Close()
}