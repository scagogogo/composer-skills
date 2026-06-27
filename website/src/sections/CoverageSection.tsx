import React from 'react'
import { useTranslation } from 'react-i18next'
import { Table, Typography, Collapse } from 'antd'
import SectionTitle from '../components/SectionTitle'

const { Title } = Typography

const CoverageSection: React.FC = () => {
  const { t } = useTranslation()

  const packagistCategories = t('coverage.packagistCategories', { returnObjects: true }) as Array<{
    category: string
    methods: string
  }>

  const composerCategories = t('coverage.composerCategories', { returnObjects: true }) as Array<{
    category: string
    count: string
    highlights: string
  }>

  const packagistColumns = [
    {
      title: t('coverage.categoryHeader'),
      dataIndex: 'category',
      key: 'category',
      width: '25%',
      render: (text: string) => <strong>{text}</strong>,
    },
    {
      title: t('coverage.methodsHeader'),
      dataIndex: 'methods',
      key: 'methods',
    },
  ]

  const composerColumns = [
    {
      title: t('coverage.categoryHeader'),
      dataIndex: 'category',
      key: 'category',
      width: '20%',
      render: (text: string) => <strong>{text}</strong>,
    },
    {
      title: t('coverage.countHeader'),
      dataIndex: 'count',
      key: 'count',
      width: '12%',
      align: 'center' as const,
      render: (text: string) => (
        <span
          style={{
            background: '#EEF2FF',
            color: '#4F46E5',
            padding: '2px 10px',
            borderRadius: 12,
            fontWeight: 600,
            fontSize: 13,
          }}
        >
          {text}
        </span>
      ),
    },
    {
      title: t('coverage.highlightsHeader'),
      dataIndex: 'highlights',
      key: 'highlights',
    },
  ]

  const collapseItems = [
    {
      key: 'packagist',
      label: <Title level={4} style={{ margin: 0 }}>{t('coverage.packagistTitle')}</Title>,
      children: (
        <Table
          columns={packagistColumns}
          dataSource={packagistCategories.map((item, i) => ({ ...item, key: i }))}
          pagination={false}
          bordered
          size="small"
        />
      ),
    },
    {
      key: 'composer',
      label: <Title level={4} style={{ margin: 0 }}>{t('coverage.composerTitle')}</Title>,
      children: (
        <Table
          columns={composerColumns}
          dataSource={composerCategories.map((item, i) => ({ ...item, key: i }))}
          pagination={false}
          bordered
          size="small"
        />
      ),
    },
  ]

  return (
    <section id="coverage" style={{ background: '#F8FAFC' }}>
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('coverage.title')} subtitle={t('coverage.subtitle')} />

        <Collapse
          defaultActiveKey={['packagist', 'composer']}
          items={collapseItems}
          style={{ background: '#fff', borderRadius: 8 }}
        />
      </div>
    </section>
  )
}

export default CoverageSection
