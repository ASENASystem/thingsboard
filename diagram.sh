#!/bin/bash
export umlFile="DIAGRAM.md"

echo "\`\`\`plantuml" > $umlFile

goplantuml -recursive . >> $umlFile

echo "\`\`\`" >> $umlFile
echo "UML diagram created"
