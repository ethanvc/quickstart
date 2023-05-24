from langchain.llms import OpenAI

# please use python 3.11 to run it, 3.9 may arise error.
llm = OpenAI(temperature=0.9)
text = "What would be a good company name for chat service?"
resp = llm(text)
print(resp)