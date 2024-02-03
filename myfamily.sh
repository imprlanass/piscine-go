curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq --argjson hero_id  "$HERO_ID" ' .[] | select( .id == $hero_id ) .connections.relatives' | tr -d '"'
