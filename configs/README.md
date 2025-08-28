# CodeRabbit Configuration Directory

This directory contains technology-specific configuration files for CodeRabbit code reviews. Each subdirectory contains a `path_instructions.yaml` file with review instructions for specific technologies.

## Structure

```
configs/
├── http/
│   └── path_instructions.yaml      # HTTP client best practices
├── kafka/
│   └── path_instructions.yaml      # Apache Kafka best practices
├── redis/
│   └── path_instructions.yaml      # Redis & caching best practices
├── scylla/
│   └── path_instructions.yaml      # ScyllaDB best practices
├── spanner/
│   └── path_instructions.yaml      # Google Spanner best practices
├── sql/
│   └── path_instructions.yaml      # PostgreSQL/CloudSQL best practices
└── temporal/
    └── path_instructions.yaml      # Temporal workflows & activities best practices
```

## File Format

Each `path_instructions.yaml` file contains an array of path instruction objects:

```yaml
- path: "**/{technology,keywords}*/**"
  instructions: |
    **Technology-specific best practices:**
    Detailed review instructions...
```

## Key Benefits

### **1. No More Duplication**
- **Before**: Each technology had duplicate instructions for file and folder patterns
- **Now**: Only define folder-based patterns, file patterns are auto-generated

### **2. Automatic Pattern Generation**
The merge script automatically creates both patterns:
- **Folder Pattern**: `**/{http,client,api,network}*/**` (matches directories)
- **File Pattern**: `**/*{http*,client*,api*,network*}*` (matches filenames)

### **3. Easy Maintenance**
- Update instructions in one place
- No need to maintain duplicate patterns
- Consistent formatting across all patterns

## Adding New Technologies

1. **Create a new directory**: `mkdir configs/newtech`
2. **Create path_instructions.yaml**:
   ```yaml
   # New Technology Best Practices
   - path: "**/{newtech,related,keywords}*/**"
     instructions: |
       **General New Technology Best Practices:**
       Your review instructions here...
   ```
3. **Run the merge tool**: `make run` or `go run cmd/merge/main.go configs/`

## Pattern Examples

| Technology | Folder Pattern | Auto-Generated File Pattern |
|------------|----------------|------------------------------|
| HTTP | `**/{http,client,api,network}*/**` | `**/*{http*,client*,api*,network*}*` |
| Kafka | `**/{kafka,producer,consumer}*/**` | `**/*{kafka*,producer*,consumer*}*` |
| Redis | `**/{redis,df,dragonfly}*/**` | `**/*{redis*,df*,dragonfly*}*` |

## Best Practices

1. **Use Descriptive Keywords**: Choose keywords that clearly identify the technology
2. **Keep Instructions Focused**: Write specific, actionable review guidance
3. **Use Consistent Formatting**: Follow the established markdown format
4. **Test Patterns**: Verify that your patterns match the intended files/directories

## Merge Process

The merge tool (`cmd/merge/main.go`) automatically:
1. Reads all `path_instructions.yaml` files
2. Generates file-based patterns from folder patterns
3. Merges everything into the final `.coderabbit.yaml`
4. Applies proper YAML formatting and indentation

## Running the Merge Tool

```bash
# Build and run
make run

# Or run directly
go run cmd/merge/main.go configs/

# Custom output file
go run cmd/merge/main.go configs/ my-coderabbit.yaml
```
