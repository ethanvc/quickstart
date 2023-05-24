from langchain.llms import OpenAI

llm = OpenAI(temperature=0.9)
text = "What would be a good company name for chat service?"
resp = llm(text)
print(resp)