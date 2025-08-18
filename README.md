# CodeRabbit Configuration Repository

A comprehensive, production-ready configuration for [CodeRabbit](https://coderabbit.ai/) that provides intelligent code review guidance for multiple technologies and best practices.

## 🚀 Features

### **Automatic Labeling**
- **Smart PR Classification**: Automatically applies relevant labels based on PR content
- **Technology-Specific Labels**: `feature`, `bugfix`, `refactor`, `documentation`, `chore`, `tests`

### **Comprehensive Technology Coverage**
- **Temporal Workflows & Activities** - Best practices for workflow orchestration
- **Redis & Caching** - Performance and security guidelines
- **HTTP Clients** - Connection pooling and timeout management
- **Apache Kafka** - Producer/consumer optimization
- **PostgreSQL/CloudSQL** - Database performance and security
- **Google Spanner** - Distributed database best practices
- **ScyllaDB** - NoSQL database optimization

### **Intelligent Path Matching**
- **Case-Insensitive**: Works regardless of naming conventions
- **Flexible Patterns**: Matches both file names and directory structures
- **Comprehensive Coverage**: Catches files like `/src/redis/main.go`, `redis_client.py`, etc.


## 🔧 Configuration Details

### **Path Pattern Strategy**

Each technology uses **dual patterns** for maximum coverage:

1. **File Pattern**: `**/*{keyword*}*` - Matches files containing keywords
2. **Directory Pattern**: `**/{keyword}*/**` - Matches files in directories with keywords

#### **Examples of What Gets Matched:**

| Technology | File Examples | Directory Examples |
|------------|---------------|-------------------|
| **Temporal** | `workflow_engine.py`, `temporal_client.go` | `/src/temporal/`, `/lib/workflow/` |
| **Redis** | `redis_cache.py`, `cache_manager.ts` | `/src/redis/`, `/lib/cache/` |
| **HTTP** | `http_client.py`, `api_service.go` | `/services/http/`, `/lib/client/` |
| **Kafka** | `kafka_producer.py`, `stream_processor.ts` | `/src/kafka/`, `/services/stream/` |
| **PostgreSQL** | `postgres_client.py`, `database_manager.go` | `/lib/postgres/`, `/db/` |
| **Spanner** | `spanner_client.py`, `spanner_service.ts` | `/src/spanner/`, `/services/spanner/` |
| **ScyllaDB** | `scylla_client.py`, `scylla_manager.go` | `/src/scylla/`, `/lib/scylla/` |


### **Adding New Technologies**

1. **Add Path Patterns**:
```yaml
- path: "**/*{your_tech*}*"
  instructions: |
    **Your Technology Best Practices:**
    # Add your specific guidelines here

- path: "**/{your_tech}*/**"
  instructions: |
    **Your Technology Best Practices:**
    # Add your specific guidelines here
```

2. **Add New Labels**:
```yaml
- label: "your_label"
  instructions: |
    Apply when the pull request meets your specific criteria.
```

### **Modifying Existing Patterns**

- **Change Keywords**: Update the glob patterns in `path` fields
- **Update Instructions**: Modify the `instructions` content for your needs
- **Adjust Thresholds**: Modify `metrics` section for performance rules

### **Performance Tuning**

```yaml
metrics:
  function_complexity:
    threshold: 15        # Adjust based on your team's standards
    enabled: true
  
  function_length:
    threshold: 150       # Adjust based on your team's standards
    enabled: true
```

## 📊 Best Practices

### **For Teams**
- **Start Small**: Begin with core technologies your team uses most
- **Iterate**: Refine patterns based on review feedback
- **Document**: Keep instructions clear and actionable
- **Review Regularly**: Update configuration as your tech stack evolves

### **For Repositories**
- **Consistent Naming**: Use clear, descriptive file and directory names
- **Technology Separation**: Organize code by technology domain
- **Documentation**: Include README files in technology-specific directories

## 🔍 Troubleshooting

### **Common Issues**

1. **Paths Not Matching**
   - Check glob pattern syntax
   - Verify case sensitivity
   - Test with sample file paths

2. **Instructions Not Triggering**
   - Ensure each `path` has `instructions`
   - Check YAML indentation
   - Validate YAML syntax

3. **Performance Issues**
   - Reduce number of path patterns
   - Simplify glob patterns
   - Use more specific keywords

### **Debugging Tips**

- **Test Patterns**: Use tools like [globster](https://globster.xyz/) to test glob patterns
- **Check Logs**: Review CodeRabbit logs for pattern matching issues
- **Validate YAML**: Use online YAML validators to check syntax

## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. **Report Issues**: Open an issue for bugs or feature requests
2. **Suggest Improvements**: Propose new technology patterns or best practices
3. **Submit PRs**: Contribute new patterns or improve existing ones
4. **Share Knowledge**: Help improve documentation and examples

### **Contribution Guidelines**

- Follow existing code style and structure
- Include clear descriptions of changes
- Test patterns with sample file paths
- Update documentation as needed

## 📚 Resources

### **Official Documentation**
- [CodeRabbit Documentation](https://docs.coderabbit.ai/)
- [YAML Configuration Guide](https://docs.coderabbit.ai/configuration)
- [Path Pattern Examples](https://docs.coderabbit.ai/configuration/path-instructions)
