package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/c12s/meridian/internal/domain"
	"github.com/c12s/meridian/internal/store"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	neo4jAddress := os.Getenv("NEO4J_ADDRESS")
	driver, err := neo4j.NewDriver(fmt.Sprintf("bolt://%s", neo4jAddress), neo4j.NoAuth())
	if err != nil {
		log.Fatal(err)
	}
	dbName := os.Getenv("NEO4J_DB_NAME")
	quotas := store.NewResourceQuotaNeo4jStore(driver, dbName)
	apps := store.NewAppNeo4jStore(driver, dbName, quotas)
	namespaces := store.NewNamespaceNeo4jStore(driver, dbName, quotas, apps)

	ns := domain.NewNamespace("c12s", "dev", "v1", map[string]string{"label1": "val1", "label2": "val2"})
	ns.AddResourceQuota("mem", 5)
	ns.AddResourceQuota("cpu", 2)
	ns.AddResourceQuota("disk", 100)
	err = namespaces.Add(ns, nil)
	if err != nil {
		log.Fatal(err)
	}

	ns, err = namespaces.Get(ns.GetId())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ns)

	ns2 := domain.NewNamespace("c12s", "test", "v1", map[string]string{"label1": "val1", "label2": "val2"})
	ns2.AddResourceQuota("mem", 2)
	// ns2.AddResourceQuota("mem", 8)
	ns2.AddResourceQuota("cpu", 1)
	ns2.AddResourceQuota("disk", 50)

	err = namespaces.Add(ns2, &ns)
	if err != nil {
		log.Fatal(err)
	}

	ns2, err = namespaces.Get(ns2.GetId())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ns2)

	ns3 := domain.NewNamespace("c12s", "prod", "v1", map[string]string{"label1": "val1", "label2": "val2"})
	ns3.AddResourceQuota("mem", 2)
	// ns2.AddResourceQuota("mem", 8)
	ns3.AddResourceQuota("cpu", 1)
	ns3.AddResourceQuota("disk", 50)

	err = namespaces.Add(ns3, &ns)
	if err != nil {
		log.Fatal(err)
	}

	ns3, err = namespaces.Get(ns3.GetId())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ns3)

	ns4 := domain.NewNamespace("c12s", "prod2", "v1", map[string]string{"label1": "val1", "label2": "val2"})
	ns4.AddResourceQuota("mem", 2)
	// ns2.AddResourceQuota("mem", 8)
	ns4.AddResourceQuota("cpu", 1)
	ns4.AddResourceQuota("disk", 50)

	err = namespaces.Add(ns4, &ns3)
	if err != nil {
		log.Fatal(err)
	}

	ns4, err = namespaces.Get(ns4.GetId())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ns4)

	app := domain.NewApp(ns4, "app", "v2")
	app.AddResourceQuota("mem", 0.5)
	app.AddResourceQuota("cpu", 0.5)
	app.AddResourceQuota("disk", 10)
	err = apps.Add(app)
	if err != nil {
		log.Fatal(err)
	}

	nsApps, err := apps.FindChildren(ns3)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(nsApps)

	tree, err := namespaces.GetHierarchy(ns.GetId())
	if err != nil {
		log.Fatal(err)
	}
	treeJson, err := json.MarshalIndent(tree.Root, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(treeJson))
}
