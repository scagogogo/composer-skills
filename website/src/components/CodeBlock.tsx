import React from 'react'
import SyntaxHighlighter from 'react-syntax-highlighter/dist/esm/prism-light'
import { vscDarkPlus } from 'react-syntax-highlighter/dist/esm/styles/prism'
import { Typography, Tooltip } from 'antd'
import { CopyOutlined, CheckOutlined } from '@ant-design/icons'

interface CodeBlockProps {
  code: string
  language?: string
}

const CodeBlock: React.FC<CodeBlockProps> = ({ code, language = 'go' }) => {
  const [copied, setCopied] = React.useState(false)

  const handleCopy = () => {
    navigator.clipboard.writeText(code).then(() => {
      setCopied(true)
      setTimeout(() => setCopied(false), 2000)
    })
  }

  return (
    <div style={{ position: 'relative', borderRadius: 8, overflow: 'hidden' }}>
      <Tooltip title={copied ? 'Copied!' : 'Copy'}>
        <Typography.Link
          onClick={handleCopy}
          style={{
            position: 'absolute',
            top: 8,
            right: 8,
            zIndex: 1,
            color: '#94A3B8',
            fontSize: 16,
          }}
        >
          {copied ? <CheckOutlined /> : <CopyOutlined />}
        </Typography.Link>
      </Tooltip>
      <SyntaxHighlighter
        language={language}
        style={vscDarkPlus}
        customStyle={{
          margin: 0,
          borderRadius: 8,
          fontSize: 13,
          padding: '16px 20px',
        }}
        showLineNumbers={false}
      >
        {code}
      </SyntaxHighlighter>
    </div>
  )
}

export default CodeBlock
