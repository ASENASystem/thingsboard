#!/bin/bash

# Create Git `pre-commit` script running this file in order to generate UML diagram
# of this project before each commit.

export umlFile="DIAGRAM.md"

echo "\`\`\`plantuml" > $umlFile

goplantuml -recursive . >> $umlFile

echo "\`\`\`" >> $umlFile
echo "UML diagram created"

git add $umlFile